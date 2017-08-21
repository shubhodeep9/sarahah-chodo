#TODO
import requests

headers = {
	'Accept'			: '*/*',
	'Accept-Encoding'	: 'gzip, deflate, br',
	'Accept-Language'	: 'en-US,en;q=0.8,hi;q=0.6',
	'Content-Type'		: 'application/x-www-form-urlencoded; charset=UTF-8',
	'Cookie'			: '.AspNetCore.Antiforgery.w5W7x28NAIs=CfDJ8MQSRebrM95Pv2f7WNJmKQXq7A-vdGN5k3C38lS0_iEwo2SCHmybZCM5x1WYbyrTrV1Gw0BRkdlzHMZH60zvB-qhDCY3e4llB1mCkBmz2hFr-toh104eg-rr-bY1xDa-ResQz962m5HnGak4InVdkTQ; ai_user=PVkSB|2017-08-20T18:48:05.938Z; .AspNetCore.Identity.Application=CfDJ8Nza2ksti91GhkboRyBeTr5EeRn0MzDNWSVqWvkxgNPOpS9HStOtBysXoMiH450q_jJLpJbYw7ahNwbbqIksOR1Mla0VseGXM-OZPOcr8l7fLtHoesWPBWHXexJKPhGjIeEAdpwk0THaJw1AkuyD9ZJFmZSS_198Tk-aM7zqwHhKz7mtZNF2JKFv4Aw2BrhysWHSbKpWCIaxb20D_x9hbHiXJLclhOKWevoNC-LKvyDwL3FrMao5pLe9Uw4f3ZJFeKDjfbuUf_Va_OPbTpDiDL5GoR8XBkDi-_Jqe9XTCdUkPgkOoodcSIbJ1ZfWw0XONwT8wKsCbYpK8TPhAanIpE2VDm_KxVxum_YiK-o7GfdHvhQFnJpyFVn-yZ64aEdOFgSMOQinTHou9X_MvgBgwf2UW8hEtAKZ7OnlE2HmxxSDKT8qSKGywWzmDW5ixbcEkAZ3PPG5ARe0PhcHmbSJoDrjONEM8lZvzFihuz4he_9qLGhdhqBXGyXet3Q0flRbZie8cE-gjisi9yET3TQKNnQ; ai_session=KKi1+|1503254886446|1503255100060.22',
	'Host'				: 'ambitionbox.sarahah.com',
	'Origin'			: 'https://ambitionbox.sarahah.com',
	'X-Requested-With'	: 'XMLHttpRequest'
}

payload = {
	'userId'						: '2fb5b368-953d-4658-acaf-425ade091f8e',
	'text'							: 'Yo',
	'captchaResponse'				: ''
}

r = requests.post('https://ambitionbox.sarahah.com/Messages/SendMessage', data=payload, headers=headers)

print(r.text)
if r.status_code == 200:
	print('Yay!')

print(r.status_code)

