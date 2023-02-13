// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/fanchunke/xgpt3/conversation/ent/schema","Package":"github.com/fanchunke/xgpt3/conversation/ent/chatent","Schemas":[{"name":"Message","config":{"Table":""},"edges":[{"name":"spouse","type":"Message","field":"spouse_id","unique":true},{"name":"session","type":"Session","field":"session_id","ref_name":"messages","unique":true,"inverse":true}],"fields":[{"name":"session_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"会话Id"},{"name":"from_user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"size":50}},"comment":"消息发送者Id"},{"name":"to_user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"size":50}},"comment":"消息接收者Id"},{"name":"content","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"消息内容"},{"name":"spouse_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"default":"CURRENT_TIMESTAMP"}}}],"indexes":[{"fields":["session_id","from_user_id","created_at"]},{"fields":["session_id","to_user_id","created_at"]}]},{"name":"Session","config":{"Table":""},"edges":[{"name":"messages","type":"Message"}],"fields":[{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"size":50}},"comment":"用户Id"},{"name":"status","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"会话是否开启"},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"default":"CURRENT_TIMESTAMP"}}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"default":"CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}}},{"name":"deleted_at","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":2,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["status","user_id"]}]}],"Features":["sql/lock","sql/upsert","privacy","entql","schema/snapshot","sql/modifier","sql/execquery"]}`
