package code

func FindErr(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return "未知错误"
}

var Msg = map[int]string{
	403:   "无效身份",
	1000:  "系统异常",
	1001:  "系统未准备就绪",
	1002:  "查询异常",
	1003:  "操作redis异常",
	1004:  "系统繁忙,请稍后再试",
	1010:  "用户不存在",
	1011:  "用户会话不存在,请重试",
	1012:  "账户不存在",
	1013:  "合约品种不存在",
	1014:  "合约不存在",
	1015:  "指数价格不存在",
	1016:  "对手价不存在",
	1017:  "查询订单不存在",
	1018:  "主账号不存在",
	1019:  "主账号不在开通子账号白名单里",
	1020:  "您的子账号数量已超出限制,请联系客服",
	1021:  "开户失败。您的主账号尚未开通合约交易权限,请前往开通",
	1030:  "输入错误",
	1031:  "非法的请求来源",
	1032:  "访问次数超出限制",
	1033:  "合约周期字段值错误",
	1034:  "报单价格类型字段值错误",
	1035:  "报单方向字段值错误",
	1036:  "报单开平字段值错误",
	1037:  "倍数不符合要求,请联系客服",
	1038:  "下单价格超出精度限制,请修改后下单",
	1039:  "买入价必须低于{0}{1},卖出价必须高于{2}{3}",
	1040:  "下单数量不能为空或者不能小于0, 请修改后下单",
	1041:  "下单数量超出限制 ({0}张),请修改后下单",
	1042:  "下单数量+挂单数量+持仓数量超过了单用户多仓持仓限制({0}张),请修改后下单",
	1043:  "下单数量+挂单数量+持仓数量超过了单用户空仓持仓限制({0}张), 请修改后下单",
	1044:  "触发平台限仓,请修改后下单",
	1045:  "当前有挂单,无法切换倍数",
	1046:  "当前合约持仓不存在",
	1047:  "可用担保资产不足",
	1048:  "可平量不足",
	1049:  "暂不支持市价开仓",
	1050:  "客户报单号重复",
	1051:  "没有可撤销的订单",
	1052:  "批量撤单、下单的订单数量超过平台限制数量",
	1053:  "无法获取最新价格区间",
	1054:  "无法获取最新价",
	1055:  "价格不合理, 下单后将导致账户权益小于0 , 请修改价格后下单",
	1056:  "结算中,暂时无法下单/撤单",
	1057:  "暂停交易中,暂时无法下单",
	1058:  "停牌中,暂时无法下单",
	1059:  "交割中,暂时无法下单/撤单",
	1060:  "合约处于非交易状态,暂时无法下单",
	1061:  "订单不存在",
	1062:  "撤单中,请耐心等待",
	1063:  "订单已成交",
	1064:  "报单主键冲突",
	1065:  "客户报单号不是整数",
	1066:  "{0}字段不能为空",
	1067:  "{0}字段不合法",
	1068:  "导出错误",
	1069:  "价格不合理",
	1070:  "数据为空,无法导出",
	1071:  "订单已撤,无法撤单",
	1072:  "卖出价必须低于{0}{1}",
	1073:  "仓位异常,请联系客服",
	1074:  "下单异常,请联系客服",
	1075:  "价格不合理, 下单后将导致担保资产率小于0 , 请修改价格后下单",
	1076:  "盘口无数据,请稍后再试",
	1077:  "交割结算中,当前品种资金查询失败",
	1078:  "交割结算中,部分品种资金查询失败",
	1079:  "交割结算中,当前品种持仓查询失败",
	1080:  "交割结算中,部分品种持仓查询失败",
	1081:  "{0}合约计划委托订单数量不得超过{1}",
	1082:  "触发类型参数错误",
	1083:  "您的仓位已进入强平接管,暂时无法下单",
	1084:  "您的合约API挂单接口被禁用,请于{0} (GMT+8) 后再试",
	1085:  "计划委托下单失败,请修改价格再次下单或联系客服",
	1086:  "{0}合约暂时限制{1}端开仓,请联系客服",
	1087:  "{0}合约暂时限制{1}端平仓,请联系客服",
	1088:  "{0}合约暂时限制{1}端撤单,请联系客服",
	1089:  "{0}账户暂时限制划转,请联系客服",
	1090:  "担保资产率小于0, 无法下单",
	1091:  "账户权益小于0, 无法下单",
	1092:  "闪电平仓取盘口第{0}档的价格, 下单后将导致账户权益小于0 , 请改为手动输入价格或使用对手价下单",
	1093:  "闪电平仓取盘口第{0}档的价格, 下单后将导致担保资产率小于0 , 请改为手动输入价格或使用对手价下单",
	1094:  "倍数不能为空, 请切换倍数或联系客服",
	1095:  "合约处于非交易状态, 暂时无法切换倍数",
	1100:  "您没有开仓权限,请联系客服",
	1101:  "您没有平仓权限,请联系客服",
	1102:  "您没有转入权限,请联系客服",
	1103:  "您没有转出权限,请联系客服",
	1104:  "合约交易受限,当前禁止交易",
	1105:  "合约交易受限,当前只能平仓",
	1106:  "合约交割结算中,暂时无法划转",
	1108:  "Dubbo调用异常",
	1109:  "子账号没有开仓权限,请联系客服",
	1110:  "子账号没有平仓权限,请联系客服",
	1111:  "子账号没有入金权限,请联系客服",
	1112:  "子账号没有出金权限,请联系客服",
	1113:  "子账号没有交易权限,请登录主账号授权",
	1114:  "子账号没有划转权限,请登录主账号授权",
	1115:  "您没有访问此子账号的权限",
	1200:  "登录失败,请重试",
	1220:  "您尚未开通合约交易,无访问权限",
	1221:  "币币账户总资产不满足合约开通条件",
	1222:  "开户天数不满足合约开通条件",
	1223:  "VIP等级不满足合约开通条件",
	1224:  "您所在的国家/地区不满足合约开通条件",
	1225:  "开通合约失败",
	1226:  "合约已开户,无法重复开户",
	1227:  "火币合约暂不支持子账户,请返回退出子账户,切换主账户登录",
	1228:  "您尚未开通合约交易, 请先开通",
	1229:  "重复同意协议",
	1230:  "您尚未做风险认证",
	1231:  "您尚未做身份认证",
	1232:  "您上传的图片格式/大小不符合要求,请重新上传",
	1233:  "您尚未开通高倍数协议 (使用高倍数请先使用主账号登录web或APP端同意高倍数协议)",
	1234:  "{0}合约未完成的开仓委托数量不得超过{1}笔",
	1235:  "{0}合约未完成的平仓委托数量不得超过{1}笔",
	1250:  "无法获取HT_token",
	1251:  "无法获取BTC净资产,请稍后再试",
	1252:  "无法获取币币账户资产,请稍后再试",
	1253:  "签名验证错误",
	1254:  "子账号无权限开通合约，请前往web端登录主账号开通",
	1300:  "划转失败",
	1301:  "可划转余额不足",
	1302:  "系统划转错误",
	1303:  "单笔转出的数量不能低于{0}{1}",
	1304:  "单笔转出的数量不能高于{0}{1}",
	1305:  "单笔转入的数量不能低于{0}{1}",
	1306:  "单笔转入的数量不能高于{0}{1}",
	1307:  "您当日累计转出量超过{0}{1}, 暂无法转出",
	1308:  "您当日累计转入量超过{0}{1}, 暂无法转入",
	1309:  "您当日累计净转出量超过{0}{1}, 暂无法转出",
	1310:  "您当日累计净转入量超过{0}{1}, 暂无法转入",
	1311:  "超过平台当日累计最大转出量限制, 暂无法转出",
	1312:  "超过平台当日累计最大转入量限制, 暂无法转入",
	1313:  "超过平台当日累计最大净转出量限制, 暂无法转出",
	1314:  "超过平台当日累计最大净转入量限制, 暂无法转入",
	1315:  "划转类型错误",
	1316:  "划转冻结失败",
	1317:  "划转解冻失败",
	1318:  "划转确认失败",
	1319:  "查询可划转金额失败",
	1320:  "此合约在非交易状态中, 无法进行系统划转",
	1321:  "划转失败, 请稍后重试或联系客服",
	1322:  "划转金额必须大于0",
	1323:  "服务异常, 划转失败, 请稍后再试",
	1325:  "设置交易单位失败",
	1326:  "获取交易单位失败",
	1327:  "无划转权限, 划转失败, 请联系客服",
	1328:  "无划转权限, 划转失败, 请联系客服",
	1329:  "无划转权限, 划转失败, 请联系客服",
	1330:  "无划转权限, 划转失败, 请联系客服",
	1331:  "超出划转精度限制(8位), 请修改后操作",
	1332:  "永续合约不存在",
	1333:  "开通跟单吃单协议失败",
	1334:  "查询跟单吃单协议失败",
	1335:  "查询跟单吃单二次确认设置失败",
	1336:  "更新跟单吃单二次确认设置失败",
	1337:  "查询跟单吃单设置失败",
	1338:  "更新跟单吃单设置失败",
	1339:  "昵称含有不合法词汇, 请修改",
	1340:  "昵称已被使用, 请修改",
	1341:  "报名阶段已结束",
	1342:  "子账号无法设置昵称",
	1343:  "指标失效, 请重新设置",
	1344:  "抱歉, 目前可最多对{0}个合约创建行情提醒",
	1345:  "抱歉, {0}合约目前可最多创建{1}个提醒",
	1346:  "该指标已存在, 请勿重复设置",
	1347:  "{0}参数错误, 请修改",
	1348:  "该合约不支持全仓模式",
	1349:  "委托单倍数与当前持仓的倍数不符, 请先切换倍数",
	1401:  "委托价必须小于行权价",
	1403:  "{0}合约止盈止损订单的委托数量不得超过{1}",
	1404:  "止盈止损订单仅支持与开仓订单绑定",
	1405:  "止盈价不得{0}{1}{2}",
	1406:  "您的抽奖次数已用完",
	1407:  "止损价不得{0}{1}{2}",
	1408:  "该止盈止损委托单未生效, 无法撤销",
	1409:  "您没有止盈止损订单权限,请联系客服",
	1410:  "批量操作子账号的数量不能超过{0}个",
	1411:  "结算中, 暂时无法查询订单信息",
	1412:  "{0}不符合价格精度限制{1}",
	1413:  "您没有跟踪委托订单权限,请联系客服",
	1414:  "您尚未开通网格交易协议(使用网格交易请先使用主账号登录web或APP端同意协议)",
	1415:  "终止价不得在网格价格范围内,请修改!",
	1416:  "超出最大运行时长({0}天{1}时),请修改!",
	1417:  "超出网格数量范围({0}~{1}个),请修改!",
	1418:  "最多同时运行{0}个网格, 请先终止其它网格",
	1419:  "超出初始化保证金范围({0}~{1}}{2})",
	1420:  "您没有合约的网格交易权限, 请联系客服",
	1421:  "当前合约有委托单或持仓, 请先撤销或平仓",
	1422:  "预计每格收益率小于0,请修改!",
	1423:  "网格最低价和最高价的价差不合理,请修改!",
	1424:  "该网格已因其它原因被终止, 无法修改或手动终止",
	1425:  "回调比例必须{0}{1}, 请修改!",
	1426:  "激活价必须{0}最新价",
	1427:  "{0}合约跟踪委托订单数量不得超过{1}",
	1428:  "相同合约的优惠券只能领取1张",
	1429:  "已领取,请勿重复领取",
	1430:  "此优惠券已失效,请刷新",
	1431:  "系统维护中,预计GMT+8 {0} 可恢复",
	1432:  "存在初始化中的网格,暂时无法下单",
	1433:  "您有限价单导致网格终止,请前往历史委托查看详情",
	1434:  "小于最小初始担保资产({0}{1}), 导致每个网格的数量小于最小下单量, 请修改!",
	1435:  "该网格已被您手动终止",
	1436:  "网格超时, 自动终止",
	1437:  "系统原因导致网格终止, 请联系客服",
	1438:  "网格触发终止条件, 已终止",
	1439:  "网格触发强平, 已终止",
	1440:  "{0}合约撤销失败",
	1441:  "触发价必须小于网格最高终止价, 大于网格最低终止价, 请修改!",
	1442:  "有效时长必须大于已运行时长1分钟以上, 请修改!",
	1443:  "合约{0}交割导致网格终止",
	1450:  "您所在的风险等级不支持使用当前倍数",
	1451:  "您所在的风险等级不支持使用当前倍数, 请登录主账号查看",
	1452:  "网格订单数量超出平台数量限制, 暂时无法下单",
	1453:  "计划委托订单数量超出平台数量限制, 暂时无法下单",
	1454:  "止盈止损订单数量超出平台数量限制, 暂时无法下单",
	1455:  "跟踪委托订单数量超出平台数量限制, 暂时无法下单",
	12001: "无效的提交时间",
	12002: "错误的签名版本",
	12003: "错误的签名方法",
	12004: "密钥已经过期",
	12005: "ip地址错误",
	12006: "提交时间不能为空",
	12007: "公钥错误",
	12008: "校验失败",
	12009: "用户被锁定或不存在",
}