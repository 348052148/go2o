/**
 * Copyright 2015 @ z3q.net.
 * name : mss_rep
 * author : jarryliu
 * date : 2015-07-27 09:03
 * description :
 * history :
 */
package repository

import (
	"github.com/jsix/gof/db"
	"go2o/src/core"
	"go2o/src/core/domain/interface/partner/mss"
	"go2o/src/core/variable"
)

var _ mss.IMssRep = new(MssRep)

type MssRep struct {
	_conn db.Connector
}

func NewMssRep(conn db.Connector) mss.IMssRep {
	return &MssRep{
		_conn: conn,
	}
}

// 获取邮箱模板
func (this *MssRep) GetMailTemplate(partnerId, id int) *mss.MailTemplate {
	var e mss.MailTemplate
	if err := this._conn.GetOrm().Get(id, &e); err == nil {
		return &e
	}
	return nil
}

// 保存邮箱模版
func (this *MssRep) SaveMailTemplate(v *mss.MailTemplate) (int, error) {
	var err error
	var orm = this._conn.GetOrm()
	if v.Id > 0 {
		_, _, err = orm.Save(v.Id, v)
	} else {
		_, _, err = orm.Save(nil, v)
		this._conn.ExecScalar("SELECT MAX(id) FROM pt_mail_template WHERE partner_id=?", &v.Id, v.PartnerId)
	}
	return v.Id, err
}

// 获取所有的邮箱模版
func (this *MssRep) GetMailTemplates(partnerId int) []*mss.MailTemplate {
	var list = []*mss.MailTemplate{}
	this._conn.GetOrm().Select(&list, "partner_id=?", partnerId)
	return list
}

// 删除邮件模板
func (this *MssRep) DeleteMailTemplate(partnerId, id int) error {
	_, err := this._conn.GetOrm().Delete(mss.MailTemplate{}, "partner_id=? AND id=?", partnerId, id)
	return err
}

// 加入到发送对列
func (this *MssRep) JoinMailTaskToQueen(v *mss.MailTask) error {
	var err error
	if v.Id > 0 {
		_, _, err = this._conn.GetOrm().Save(v.Id, v)
	} else {
		_, _, err = this._conn.GetOrm().Save(nil, v)
		if err == nil {
			err = this._conn.ExecScalar("SELECT max(id) FROM pt_mail_queue", &v.Id)
		}
	}

	if err == nil {
		rc := core.GetRedisConn()
		defer rc.Close()
		rc.Do("RPUSH", variable.KvNewMailTask, v.Id) // push to queue
	}
	return err
}
