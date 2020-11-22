package restful

import (
	"github.com/hoisie/web"
	"github.com/sirupsen/logrus"
)

func init_router(server *web.Server) {
	logrus.Info("Инициализация роутов...")

	/*
		// auth

		server.Get(`/api/auth`, handler_get_info_by_token)
		server.Post(`/api/auth`, handler_login)
		server.Delete(`/api/auth`, handler_logout)

		// // devices

		server.Get(`/api/devices/([a-z|\_|\-]*)`, handler_devices_by_type)
		server.Get("/api/device/([0-9]*)", handler_get_device)
		server.Get(`/api/device/([0-9]*)/connected/([a-z|\_|\-]*)`, handler_get_device_connected)
		server.Get(`/api/device/([0-9]*)/parameters`, handler_get_device_parameters)
		server.Put("/api/device/([0-9]*)", handler_put_device)
		server.Post("/api/device/([0-9]*)", handler_put_device)
		server.Delete("/api/device/([0-9]*)", handler_delete_device)

		server.Get(`/api/device/([0-9]*)/device2device`,
			handler_devices_get_device2device)
		server.Put(`/api/device2device/([0-9]*)`,
			handler_devices_set_device2device)
		server.Post(`/api/device2device/([0-9]*)`,
			handler_devices_set_device2device)
		server.Delete(`/api/device2device/([0-9]*)`,
			handler_devices_delete_device2device)

		// marks

		server.Get(`/api/userlog`, handler_userlog)

		server.Get(`/api/marks/([a-z|\_]*)`, handler_marks_by_type)
		server.Post(`/api/mark/([0-9]*)`, handler_edit_mark)
		server.Put(`/api/mark/([0-9]*)`, handler_edit_mark)
		server.Delete(`/api/mark/([0-9]*)`, handler_delete_mark)

		// server.Get(`/api/parameters`, handler_parameters)
		// server.Post(`/api/parameter/([0-9]*)`, handler_edit_parameter)
		// server.Put(`/api/parameter/([0-9]*)`, handler_edit_parameter)
		// server.Delete(`/api/parameter/([0-9]*)`, handler_delete_parameter)

		// models

		server.Get(`/api/models/([a-z|\_]*)/([0-9]*)`, handler_models_by_type_by_mark)
		server.Post(`/api/model/([0-9]*)`, handler_edit_model)
		server.Put(`/api/model/([0-9]*)`, handler_edit_model)
		server.Delete(`/api/model/([0-9]*)`, handler_delete_model)

		// payers

		server.Get(`/api/payers`, handler_get_payers)
		server.Put(`/api/payer/([0-9]*)`, handler_edit_payer)
		server.Post(`/api/payer/([0-9]*)`, handler_edit_payer)
		server.Delete(`/api/payer/([0-9]*)`, handler_delete_payer)

		// place

		// Получаем самый корневой для пользователя объект

		server.Get(`/api/places`, handler_get_places)

		// все дочерние эдементы на родительском
		server.Get(`/api/places/0`, handler_get_places)

		// Получение всех дочерних элементов на указанном
		server.Get(`/api/places/([0-9]*)`, handler_get_places_by_id)
		server.Put(`/api/place/([0-9]*)`, handler_add_place)
		server.Get(`/api/place/([0-9]*)/connected/([a-z|\_]*)`, handler_get_place_connected_device)
		server.Post(`/api/place/([0-9]*)`, handler_add_place)

		server.Post(`/api/place/([0-9]*)/connected/device`, handler_add_place_connected_device)
		server.Delete(`/api/place/([0-9]*)/connected/device`, handler_delete_place_connected_device)

		server.Delete(`/api/place/([0-9]*)`, handler_delete_place)

		// Личные кабинеты
		// Получение информации о главном

		server.Get(`/api/personal_area`, handler_get_personal_area)
		server.Get(`/api/personal_area/([0-9]*)`, handler_get_personal_area_children)
		server.Put(`/api/personal_area/([0-9]*)`, handler_edit_personal_area)
		server.Post(`/api/personal_area/([0-9]*)`, handler_edit_personal_area)
		server.Delete(`/api/personal_area/([0-9]*)`, handler_delete_personal_area)

		// Пользователи

		server.Get(`/api/users`, handler_get_users)
		server.Get(`/api/user/([0-9]*)`, handler_get_user)
		server.Put(`/api/user/([0-9]*)`, handler_edit_user)
		server.Post(`/api/user/([0-9]*)`, handler_edit_user)
		server.Delete(`/api/user/([0-9]*)`, handler_delete_user)

		// Данные

		server.Get(`/api/data/device/([0-9]*)/last_data`, handler_get_last_data_device)
		server.Post(`/api/device/([0-9]*)/fast_pool`, handler_device_fast_pool)

		// /api/data/device/10/range/now/20201012131512/2020101613
		server.Get(`/api/data/device/([0-9]*)/([0-9]{10})/([0-9]{10})`,
			handler_get_data_device_range)
		server.Get(`/api/data/device/([0-9]*)/([0-9]{8})/([0-9]{8})`,
			handler_get_data_device_range)
		server.Get(`/api/data/device/([0-9]*)/([0-9]{6})/([0-9]{6})`,
			handler_get_data_device_range)
		server.Get(`/api/data/device/([0-9]*)/([0-9]{4})/([0-9]{4})`,
			handler_get_data_device_range)

		server.Get(`/api/report/device/([0-9]*)/([0-9]{10})/([0-9]{10})/([a-z|0-9]*)`,
			handler_get_report_device_range)
		server.Get(`/api/report/device/([0-9]*)/([0-9]{8})/([0-9]{8})/([a-z|0-9]*)`,
			handler_get_report_device_range)
		server.Get(`/api/report/device/([0-9]*)/([0-9]{6})/([0-9]{6})/([a-z|0-9]*)`,
			handler_get_report_device_range)
		server.Get(`/api/report/device/([0-9]*)/([0-9]{4})/([0-9]{4})/([a-z|0-9]*)`,
			handler_get_report_device_range)

		server.Get("/api/rules", handler_get_rules)

		server.Get("/api/roles", handler_get_roles)
		server.Post("/api/role/([0-9]*)", handler_post_role)
		server.Put("/api/role/([0-9]*)", handler_post_role)
		server.Delete("/api/role/([0-9]*)", handler_delete_role)

		server.Get(`/api/([a-z|\_]*)/([0-9]*)/object_rules`,
			handler_get_object_rule_by_object)
		server.Get("/api/user/([0-9]*)/object_rules",
			handler_get_object_rule_by_user)
		server.Post("/api/object_rule/([0-9]*)",
			handler_post_object_rule)
		server.Put("/api/object_rule/([0-9]*)",
			handler_post_object_rule)
		server.Delete("/api/object_rule/([0-9]*)",
			handler_delete_object_rule)

		server.Get("/ws", handler_websocket)

		server.Get("/version", handler_version)

		server.Get("/api/timezones", handler_timezones)
		server.Get("/api/poolers", handler_poolers_all)
		server.Get("/api/poolers/(.*)", handler_poolers_by_uid)

		server.Match("DELETE", `(.*)`, handler_option)
		server.Match("OPTIONS", `(.*)`, handler_option)

		go deamon_actuality_ws_pooler()
	*/

}
