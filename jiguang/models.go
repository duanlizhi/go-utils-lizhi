/**
 * <p>Description: (极光推送参数封装) </>
 * @author lizhi_duan
 * @date 2018/9/7 09:40
 * @version 1.0
 */
package jiguang

import (
	"errors"
)

/**
 * <p>Description: (推送对象) </p>
 * @author lizhi_duan
 * @date 2018/9/7 10:21
 */
type PushObject struct {
	Platform     interface{}   `json:"platform,omitempty"`     //必填参数，推送平台设置
	Audience     interface{}   `json:"audience,omitempty"`     //必填参数，推送设备指定
	Notification *Notification `json:"notification,omitempty"` //(可选)通知内容体。是被推送到客户端的内容。与 message 一起二者必须有其一，可以二者并存
	Message      *Message      `json:"message,omitempty"`      //(可选)消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
	SmsMessage   *SmsMessage   `json:"sms_message,omitempty"`  //(可选)短信渠道补充送达内容体
	Options      *Options      `json:"options,omitempty"`      //(可选)推送参数
	Cid          string        `json:"cid,omitempty"`          //可选，用于防止api调用端重试造成服务端的重复推送而定义
}

/**
 * <p>Description: (创建极光推送json对象) </p>
 * 默认平台为推送给所有，推送环境默认为开发环境，可自行定义
 * @author lizhi_duan
 * @date 2018/9/9 21:39
 */
func NewPushObject(prod bool) *PushObject {
	prod_default := false
	prod_default = prod
	return &PushObject{
		Platform: "all",
		Options: &Options{
			TimeToLive:     5 * 24 * 60 * 60,
			ApnsProduction: prod_default,
		},
	}
}

/**
 * <p>Description: (发送给所有的平台) </p>
 * @author lizhi_duan
 * @date 2018/9/15 10:46
 */
func (pushObj *PushObject) SetAllPlatform() {
	pushObj.Platform = "all"
}
func (pushObj *PushObject) SetPlatform(platform *Platform) error {
	if platform == nil || len(platform.osArr) == 0 {
		return errors.New("uncatch nil platform or osArr is empty")
	}
	pushObj.Platform = platform
	return nil
}

/**
 * <p>Description: (推送到所有的目标) </p>
 * @author lizhi_duan
 * @date 2018/9/15 12:22
 */
func (pushObj *PushObject) SetAllAudience() {
	pushObj.Audience = "all"
}

/**
 * <p>Description: (发送平台) </p>
 * @author lizhi_duan
 * @date 2018/9/15 10:58
 */
type Platform struct {
	osArr []string
}

/**
 * <p>Description: (创建platform对象) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:03
 */
func NewPlatform() *Platform {
	return &Platform{
		osArr: make([]string, 3),
	}
}

/**
 * <p>Description: (添加platform) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:04
 */
func (platform *Platform) AddPlatform(os string) {
	for _, v := range platform.osArr {
		if v == os {
			return
		}
	}
	switch os {
	case PLATFORM_ANDROID:
		platform.osArr = append(platform.osArr, os)
	case PLATFORM_IOS:
		platform.osArr = append(platform.osArr, os)
	case PLATFORM_WINPHONE:
		platform.osArr = append(platform.osArr, os)
	}
}

/**
 * <p>Description: (推送设备指定（必填）) </p>
 * @author lizhi_duan
 * @date 2018/9/7 10:24
 */
type Audience struct {
	Tag            []string `json:"tag,omitempty"`             //用标签来进行大规模的设备属性、用户属性分群。 一次推送最多 20 个。
	TagAnd         []string `json:"tag_and,omitempty"`         //注意与 tag 区分。一次推送最多 20 个。
	TagNot         []string `json:"tag_not,omitempty"`         //一次推送最多 20 个。
	Alias          []string `json:"alias,omitempty"`           //用别名来标识一个用户。一个设备只能绑定一个别名，但多个设备可以绑定同一个别名。一次推送最多 1000
	RegistrationId []string `json:"registration_id,omitempty"` //设备标识。一次推送最多 1000 个。客户端集成 SDK 后可获取到该值。
	Segment        []string `json:"segment,omitempty"`         //目前限制是一次只能推送一个。
	Abtest         []string `json:"abtest,omitempty"`          //目前限制一次只能推送一个。
}

/**
 * <p>Description: (添加tags) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddTag(tag ...string) {
	tags := make([]string, 0)
	for _, v := range tag {
		isExist := false
		for _, tag_ := range audience.Tag {
			if v == tag_ {
				isExist = true
				continue
			}
		}
		if !isExist {
			tags = append(tags, v)
		}
	}
	audience.Tag = append(audience.Tag, tags...)
}

/**
 * <p>Description: (添加tagAnd) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddTagAnd(tagAnd ...string) {
	var tagAndSlice []string
	for _, v := range tagAnd {
		isExist := false
		for _, tagA := range audience.TagAnd {
			if v == tagA {
				isExist = true
				continue
			}
		}
		if !isExist {
			tagAndSlice = append(tagAndSlice, v)
		}
	}
	audience.TagAnd = append(audience.TagAnd, tagAndSlice...)
}

/**
 * <p>Description: (添加tagNot) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddTagNot(tagNot ...string) {
	var tagNotSlice []string
	for _, v := range tagNot {
		isExist := false
		for _, not := range audience.TagNot {
			if v == not {
				isExist = true
				continue
			}
		}
		if !isExist {
			tagNotSlice = append(tagNotSlice, v)
		}
	}
	audience.TagNot = append(audience.TagNot, tagNotSlice...)
}

/**
 * <p>Description: (添加alias) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddAlias(alias ...string) {
	var aliasSlice []string
	for _, v := range alias {
		isExist := false
		for _, al := range audience.Alias {
			if v == al {
				isExist = true
				continue
			}
		}
		if !isExist {
			aliasSlice = append(aliasSlice, v)
		}
	}
	audience.Alias = append(audience.Alias, aliasSlice...)
}

/**
 * <p>Description: (添加registrationId) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddRegistrationId(registrationId ...string) {
	var registrationSlice []string
	for _, v := range registrationId {
		isExist := false
		for _, re := range audience.RegistrationId {
			if v == re {
				isExist = true
				continue
			}
		}
		if !isExist {
			registrationSlice = append(registrationSlice, v)
		}
	}
	audience.RegistrationId = append(audience.RegistrationId, registrationSlice...)
}

/**
 * <p>Description: (添加segment) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddSegment(segment ...string) {
	var segmentSlice []string
	for _, v := range segment {
		isExist := false
		for _, se := range audience.Segment {
			if v == se {
				isExist = true
				continue
			}
		}
		if !isExist {
			segmentSlice = append(segmentSlice, v)
		}
	}
	audience.Segment = append(audience.Segment, segmentSlice...)
}

/**
 * <p>Description: (添加Abtest) </p>
 * @author lizhi_duan
 * @date 2018/9/15 11:23
 */
func (audience *Audience) AddAbtest(abtest ...string) {
	var abtestSlice []string
	for _, v := range abtest {
		isExist := false
		for _, ab := range audience.Abtest {
			if v == ab {
				isExist = true
				continue
			}
		}
		if !isExist {
			abtestSlice = append(abtestSlice, v)
		}
	}
	audience.Abtest = append(audience.Abtest, abtestSlice...)
}

/**
 * <p>Description: (通知内容体（可选）) </p>
 * @author lizhi_duan
 * @date 2018/9/7 15:52
 */
type Notification struct {
	Alert                string                `json:"alert,omitempty"` //通知的内容在各个平台上，都可能只有这一个最基本的属性 "alert"。
	AndroidNotification  *AndroidNotification  `json:"android,omitempty"`
	IosNotification      *IosNotification      `json:"ios,omitempty"`
	WinphoneNotification *WinphoneNotification `json:"winphone,omitempty"`
}

/**
 * <p>Description: (创建消息内容体) </p>
 * @author lizhi_duan
 * @date 2018/9/9 22:06
 */
func NewNotification(alert string) *Notification {
	return &Notification{
		Alert: alert,
		AndroidNotification: &AndroidNotification{
			Alert: alert,
		},
		IosNotification: &IosNotification{
			Alert: alert,
		},
		WinphoneNotification: &WinphoneNotification{
			Alert: alert,
		},
	}
}

/**
 * <p>Description: (Android 平台上的通知，JPush SDK 按照一定的通知栏样式展示。) </p>
 * @author lizhi_duan
 * @date 2018/9/7 16:27
 */
type AndroidNotification struct {
	//通知内容(必填)
	Alert string `json:"alert"` //这里指定了，则会覆盖上级统一指定的 alert 信息；内容可以为空字符串，则表示不展示到通知栏。
	//通知标题(可选)
	Title string `json:"title,omitempty"` //如果指定了，则通知里原来展示 App 名称的地方，将展示成这个字段。
	//通知栏样式 ID(可选)
	BuilderId uint64 `json:"builder_id,omitempty"` //Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式。
	//通知栏展示优先级(可选)
	Priority int64 `json:"priority,omitempty"` //默认为 0，范围为 -2～2 ，其他值将会被忽略而采用默认。
	//通知栏条目过滤或排序(可选)
	Category string `json:"category,omitempty"` //完全依赖 rom 厂商对 category 的处理策略
	//通知栏样式类型(可选)
	Style int64 `json:"style,omitempty"` //默认为 0，还有 1，2，3 可选，用来指定选择哪种通知栏样式，其他值无效。有三种可选分别为 bigText=1，Inbox=2，bigPicture=3。
	//通知提醒方式(可选)
	AlertType int64 `json:"alert_type,omitempty"` //可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1, Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理。
	//大文本通知栏样式(可选)
	BigText string `json:"big_text,omitempty"` //当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来。支持 api 16 以上的 rom
	//文本条目通知栏样式(可选)
	Inbox map[string]string `json:"inbox,omitempty"` //当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示。支持 api 16 以上的 rom。
	//大图片通知栏样式(可选)
	BigPicPath string `json:"big_pic_path,omitempty"` //当 style = 3 时可用，可以是网络图片 url，或本地图片的 path，目前支持 .jpg 和 .png 后缀的图片。图片内容会被通知栏以大图片的形式展示出来。如果是 http／https 的 url，会自动下载；如果要指定开发者准备的本地图片就填 sdcard 的相对路径。支持 api 16 以上的 rom。
	//扩展字段(可选)
	Extras map[string]interface{} `json:"extras,omitempty"` //这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
}

/**
 * <p>Description: (iOS 平台上 APNs 通知结构。) </p>
 * @author lizhi_duan
 * @date 2018/9/7 16:39
 */
type IosNotification struct {
	//通知内容（必填）
	Alert string `json:"alert"` //这里指定内容将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。支持字符串形式也支持官方定义的 alert payload 结构，在该结构中包含 title 和 subtitle 等官方支持的 key
	//通知提示声音(可选)
	Sound string `json:"sound,omitempty"` //如果无此字段，则此消息无声音提示；有此字段，如果找到了指定的声音就播放该声音，否则播放默认声音，
	//应用角标(可选)
	Badge int64 `json:"badge,omitempty"` //如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。JPush 官方 SDK 会默认填充 badge 值为 "+1",
	//推送唤醒(可选)
	ContentAvailable bool `json:"content-available,omitempty"` //推送的时候携带 "content-available":true 说明是 Background Remote Notification，如果不携带此字段则是普通的 Remote Notification
	//通知扩展(可选)
	MutableContent bool `json:"mutable-content,omitempty"` //推送的时候携带 ”mutable-content":true 说明是支持iOS10的UNNotificationServiceExtension，如果不携带此字段则是普通的 Remote Notification。
	//可选
	Category string `json:"category,omitempty"` //IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	//附加字段(可选)
	Extras map[string]interface{} `json:"extras,omitempty"` //这里自定义 Key / value 信息，以供业务使用。
}

/**
 * <p>Description: (Windows Phone 平台上的通知。) </p>
 * @author lizhi_duan
 * @date 2018/9/7 16:46
 */
type WinphoneNotification struct {
	//通知内容(必填)
	Alert string `json:"alert"` //会填充到 toast 类型 text2 字段上。这里指定了，将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。
	//通知标题(可选)
	Title string `json:"title,omitempty"` //会填充到 toast 类型 text1 字段上。
	//点击打开的页面名称(可选)
	OpenPage string `json:"_open_pag,omitempty"` //点击打开的页面。会填充到推送信息的 param 字段上，表示由哪个 App 页面打开该通知。可不填，则由默认的首页打开。
	//扩展字段(可选)
	Extras map[string]interface{} `json:"extras,omitempty"` //作为参数附加到上述打开页面的后边。
}

/**
 * <p>Description: (消息内容体（可选），自定义消息) </p>
 * @author lizhi_duan
 * @date 2018/9/7 15:57
 */
type Message struct {
	MsgContent  string            `json:"msg_content"`            //(必填)消息内容本身
	Title       string            `json:"title,omitempty"`        //(可选)消息标题
	ContentType string            `json:"content_type,omitempty"` //(可选)消息内容类型
	Extras      map[string]string `json:"extras,omitempty"`       //(可选)JSON 格式的可选参数
}

/**
 * <p>Description: (短信补充) </p>
 * @author lizhi_duan
 * @date 2018/9/7 16:16
 */
type SmsMessage struct {
	TempId    string            `json:"temp_id"`              //(必填)短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara  map[string]string `json:"temp_para,omitempty"`  //(可选)短信模板中的参数。
	DelayTime uint64            `json:"delay_time,omitempty"` //(必填)单位为秒，不能超过 24 小时。设置为 0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
}

/**
 * <p>Description: (可选参数) </p>
 * @author lizhi_duan
 * @date 2018/9/7 16:16
 */
type Options struct {
	/**推送序号(可选)*/
	Sendno uint64 `json:"sendno,omitempty"` //纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回。值为 0 表示该 messageid 无 sendno，所以字段取值范围为非 0 的 int.
	/**离线消息保留时长(秒)（可选）*/
	TimeToLive uint64 `json:"time_to_live,omitempty"` //推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天），最长 10 天。设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到。该字段对 iOS 的 Notification 消息无效。
	/**要覆盖的消息 ID(可选)*/
	OverrideMsgId string `json:"override_msg_id,omitempty"` //如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果，即：1）该 msg_id 离线收到的消息是覆盖后的内容；2）即使该 msg_id Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知，该字段仅对 Android 有效。
	/**APNs 是否生产环境（可选）*/
	ApnsProduction bool `json:"apns_production"` //True 表示推送生产环境，False 表示要推送开发环境；如果不指定则为推送生产环境。但注意，JPush 服务端 SDK 默认设置为推送 “开发环境”。该字段仅对 iOS 的 Notification 有效。
	/**更新 iOS 通知的标识符（可选）*/
	ApnsCollapseId string `json:"apns_collapse_id,omitempty"` //APNs 新通知如果匹配到当前通知中心有相同 apns-collapse-id 字段的通知，则会用新通知内容来更新它，并使其置于通知中心首位。collapse id 长度不可超过 64 bytes。
	/**定速推送时长(分钟)(可选)*/
	BigPushDuration uint64 `json:"big_push_duration,omitempty"` //又名缓慢推送，把原本尽可能快的推送速度，降低下来，给定的 n 分钟内，均匀地向这次推送的目标用户推送。最大值为 1400。未设置则不是定速推送。
}
