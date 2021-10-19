package objectRepository

import (
	"errors"
	"net"

	"gorm.io/gorm"

	svc "sample-service"

	svcErrors "sample-service/implementation/errors"
	"sample-service/utils/netaddr"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) svc.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(a svc.Object) (uint64, error) {
	err := repo.db.Table("Object").Save(&a).Error
	return a.ID, err
}

func (repo *repository) Update(a svc.Object) error {
	err := repo.db.Table("Object").Where("id = ?", a.ID).Updates(
		map[string]interface{}{
			"ip":   a.IP,
			"meta": a.Meta,
		}).Error
	return err
}

func (repo *repository) SetMeta(a svc.Object) error {
	err := repo.db.Table("Object").Where("id = ?", a.ID).Updates(
		map[string]interface{}{
			"meta":     a.Meta,
			"has_meta": 1,
		}).Error
	return err
}

func (repo *repository) Get(id uint64) (response svc.Object, err error) {

	err = repo.db.Table("Object").First(&response, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, svcErrors.NewAppErrorWithType(svcErrors.NotFound)
		}
		return response, svcErrors.NewAppErrorWithType(svcErrors.RepositoryError)
	}

	return response, nil
}

func (repo *repository) GetTargets(id uint64) (list []map[string]interface{}, count int64, err error) {
	subQuery := repo.db.
		Select("ip").
		Table("Object").
		Where("parent_id = ?", id).
		Group("ip")

	err = repo.db.Table("(?) as tmp", subQuery).Count(&count).Error

	if err != nil {
		return list, count, svcErrors.NewAppErrorWithType(svcErrors.RepositoryError)
	}

	err = repo.db.
		Table("Object").
		Select("ip, max(peak) as peak").
		Where("parent_id = ?", id).
		Group("ip").
		Order("peak DESC").
		Limit(10).
		Find(&list).
		Error
	if err != nil {
		return list, count, svcErrors.NewAppErrorWithType(svcErrors.RepositoryError)
	}
	return list, count, nil
}

func (repo *repository) getOrWhereByIp(ip []string) (result *gorm.DB, err error) {
	orWhere := repo.db
	for _, cidr := range ip {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return result, err
		}
		network := netaddr.NetworkAddr(ipnet)
		broadcast := netaddr.BroadcastAddr(ipnet)

		if network.Equal(broadcast) {
			orWhere = orWhere.Or("ip = ?", netaddr.Ip2int(network))
		} else {
			orWhere = orWhere.Or("ip >= ? AND ip <= ?", netaddr.Ip2int(network), netaddr.Ip2int(broadcast))
		}
	}
	return orWhere, nil
}

func (repo *repository) UnderObject(as []uint32, ip []string) (count int64, err error) {
	orWhere, err := repo.getOrWhereByIp(ip)
	if err != nil {
		return count, err
	}
	err = repo.db.
		Table("Object").
		Where("`as` IN ?", as).
		Where(orWhere).
		Where("end_time = 0").
		Count(&count).
		Error

	if err != nil {
		return count, svcErrors.NewAppErrorWithType(svcErrors.RepositoryError)
	}
	return count, nil
}
