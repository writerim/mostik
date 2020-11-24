package repositories

import (
	"entity"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func Init(db *gorm.DB) {

	logrus.Info("Плательщики")
	db.AutoMigrate(entity.Payer{})

	logrus.Info("Параметры")
	db.AutoMigrate(entity.Parameter{})

	logrus.Info("Тонкие права")
	db.AutoMigrate(entity.ObjectRule{})

	logrus.Info("Параметры моделей")
	db.AutoMigrate(entity.ModelParameter{})

	logrus.Info("Конфигурация моделей")
	db.AutoMigrate(entity.ModelConfig{})

	logrus.Info("Модели устройств")
	db.AutoMigrate(entity.Model{})

	logrus.Info("Модели устройств")
	db.AutoMigrate(entity.Mark{})

	logrus.Info("Последние данные")
	db.AutoMigrate(entity.LastData{})

	logrus.Info("Связь счетчиков")
	db.AutoMigrate(entity.Device2Device{})

	logrus.Info("Устройства")
	db.AutoMigrate(entity.Device{})

	logrus.Info("Данные")
	db.AutoMigrate(entity.Data{})

	logrus.Info("Лог пользователя")
	db.AutoMigrate(&entity.UserLog{})

	logrus.Info("Пользователи")
	db.AutoMigrate(&entity.User{})

	logrus.Info("Временные зоны")
	db.AutoMigrate(&entity.TimeZone{})

	logrus.Info("Роли")
	db.AutoMigrate(entity.Role{})

	logrus.Info("Праила")
	db.AutoMigrate(&entity.Rule{})

	logrus.Info("Связь ролей и правил")
	db.AutoMigrate(&entity.RuleRoles{})

	logrus.Info("Типы объектов")
	db.AutoMigrate(entity.PlaceType{})

	logrus.Info("Объекты")
	db.AutoMigrate(entity.Place{})

	// END MIGRATE

	tzs := []entity.TimeZone{
		entity.TimeZone{
			DiffHour: -11,
			Ident:    "Pacific/Midway",
		},
		entity.TimeZone{
			DiffHour: -10,
			Ident:    "Pacific/Honolulu",
		},
		entity.TimeZone{
			DiffHour: -9,
			Ident:    "America/Anchorage",
		},
		entity.TimeZone{
			DiffHour: -8,
			Ident:    "America/Tijuana",
		},
		entity.TimeZone{
			DiffHour: -7,
			Ident:    "America/Mazatlan",
		},
		entity.TimeZone{
			DiffHour: -6,
			Ident:    "America/Chicago",
		},
		entity.TimeZone{
			DiffHour: -5,
			Ident:    "America/New_York",
		},
		entity.TimeZone{
			DiffHour: -4,
			Ident:    "America/Halifax",
		},
		entity.TimeZone{
			DiffHour: -3,
			Ident:    "America/Godthab",
		},
		entity.TimeZone{
			DiffHour: -2,
			Ident:    "America/Noronha",
		},
		entity.TimeZone{
			DiffHour: -1,
			Ident:    "Atlantic/Azores",
		},
		entity.TimeZone{
			DiffHour: 0,
			Ident:    "Europe/London",
		},
		entity.TimeZone{
			DiffHour: 1,
			Ident:    "Europe/Berlin",
		},
		entity.TimeZone{
			DiffHour: 2,
			Ident:    "Europe/Helsinki",
		},
		entity.TimeZone{
			DiffHour: 3,
			Ident:    "Europe/Moscow",
		},
		entity.TimeZone{
			DiffHour: 4,
			Ident:    "Asia/Tbilisi",
		},
		entity.TimeZone{
			DiffHour: 5,
			Ident:    "Asia/Yekaterinburg",
		},
		entity.TimeZone{
			DiffHour: 6,
			Ident:    "Asia/Almaty",
		},
		entity.TimeZone{
			DiffHour: 7,
			Ident:    "Asia/Bangkok",
		},
		entity.TimeZone{
			DiffHour: 8,
			Ident:    "Asia/Irkutsk",
		},
		entity.TimeZone{
			DiffHour: 9,
			Ident:    "Asia/Tokyo",
		},
		entity.TimeZone{
			DiffHour: 10,
			Ident:    "Asia/Vladivostok",
		},
		entity.TimeZone{
			DiffHour: 11,
			Ident:    "Australia/Sydney",
		},
		entity.TimeZone{
			DiffHour: 12,
			Ident:    "Asia/Kamchatka",
		},
		entity.TimeZone{
			DiffHour: 13,
			Ident:    "Pacific/Tongatapu",
		},
	}

	for i, item := range tzs {
		item.Id = i + 1
		item.DiffMinutes = item.DiffHour * 60

		z := ""
		if item.DiffHour > 0 {
			z = "+"
		}

		item.Title = fmt.Sprintf(`%s%02d:00 UTC`, z, item.DiffHour)

		tz := NewMysqlTimeZoneRepository(db)
		_, err := tz.Save(item)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}

	roles := []entity.Role{
		entity.Role{
			Id:             1,
			Title:          "Администратор системы",
			PersonalAreaId: 1,
		},
	}

	for _, item := range roles {
		role := NewMysqlRoleRepository(db)
		_, err := role.Save(item)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}

	roles_action := []string{
		"view",
		"delete",
		"add",
	}
	roles_objects := []string{
		"payer",
		"place",
		"user",
		"device",
		"personal_area",
		"device2device",
		"role",
	}
	index := 0

	for _, ra := range roles_action {

		for _, ro := range roles_objects {
			rule := NewMysqlRuleRepository(db)
			index += 1
			role := entity.Rule{
				Title:  fmt.Sprintf(`%s_%s`, ra, ro),
				Object: ro,
				Action: ra,
				Id:     index,
			}
			rule.Save(role)
		}

	}

	rules := NewMysqlRuleRepository(db)
	all_rules, err := rules.GetAll()
	if err != nil {
		logrus.Warn(err.Error())
	}
	for i, item := range all_rules {
		rr := entity.RuleRoles{
			Id:     i + 1,
			RoleId: 1,
			RuleId: item.Id,
		}
		rrr := NewMysqlRuleRoleRepository(db)
		_, err := rrr.Save(rr)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}

	place_types := []entity.PlaceType{
		entity.PlaceType{
			Title: "Город",
			Ident: "gorod",
		},
		entity.PlaceType{
			Title: "Район",
			Ident: "raion",
		},
		entity.PlaceType{
			Title: "Улица",
			Ident: "ulica",
		},
		entity.PlaceType{
			Title: "Жилое здание",
			Ident: "zhiloe",
		},
		entity.PlaceType{
			Title: "Офисное здание",
			Ident: "ofisnoe",
		},
		entity.PlaceType{
			Title: "Промышленное здание",
			Ident: "prom",
		},
		entity.PlaceType{
			Title: "Строение",
			Ident: "stroenie",
		},
		entity.PlaceType{
			Title: "Торговое здание",
			Ident: "torgovoe",
		},
		entity.PlaceType{
			Title: "Корпус",
			Ident: "corpus",
		},
		entity.PlaceType{
			Title: "Подъезд",
			Ident: "podyezd",
		},
		entity.PlaceType{
			Title: "Квартира",
			Ident: "kvartira",
		},
		entity.PlaceType{
			Title: "Офис",
			Ident: "ofis",
		},
		entity.PlaceType{
			Title: "Личный кабинет",
			Ident: "lk",
		},
		entity.PlaceType{
			Title: "Системный",
			Ident: "system",
		},
	}

	for i, item := range place_types {
		item.Id = i + 1
		pt := NewMysqlPlaceTypeRepository(db)
		_, err := pt.Save(item)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}

	personal_areas := []entity.PersonalArea{
		entity.PersonalArea{
			Title:    "Система",
			ParentId: 0,
			UserId:   1,
			Location: "UTC",
		},
	}

	for i, item := range personal_areas {
		item.Id = i + 1
		pa := NewMysqlPersonalAreaRepository(db)
		_, err := pa.Save(item)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}

	user := NewMysqlUserRepository(db)

	if count, err := user.GetCountAll(); err == nil && count == 0 {

		users := []entity.User{entity.User{
			Login:    "admin",
			Password: "admin",
			Email:    "example@mail.ru",
		}}

		for i, item := range users {
			item.Id = i + 1
			item.PersonalAreaId = 1

			user := NewMysqlUserRepository(db)
			_, err := user.Save(item)
			if err != nil {
				logrus.Warn(err.Error())
			}
		}
	}

}
