:: id 5415587
:: api_key UPKOHCZhfTNMEI9B1Xnf5PRK
:: secret_key GYHsMLE820jkw1cBGnNbIzoeVVxdDUfs
:: 
:: curl "https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=UPKOHCZhfTNMEI9B1Xnf5PRK&client_secret=GYHsMLE820jkw1cBGnNbIzoeVVxdDUfs"
:: 响应：
:: {
::   "access_token":"24.774d214d120fb13f55b05b679004fab4.2592000.1427454902.282335-5415587",
::   "session_key":"9mzdDAJyBP7mE\/7iSyzd3eAMmOB3i\/H4JHJgK8MN4RO+RKEIX+MS\/tCmx0K4o7Tlwle8wzL9\/hvaj7LygMfVIF3zLBt2",
::   "scope":"public wise_adapt lebo_resource_base lightservice_public",
::   "refresh_token":"25.bf21f1da203b27e2da50549407f911bd.315360000.1740222902.282335-5415587",
::   "session_secret":"b7825996cb6b0724da5af9ca8323e90b",
::   "expires_in":2592000
:: }
:: 
:: token一个月有效 24.774d214d120fb13f55b05b679004fab4.2592000.1427454902.282335-5415587
:: 

curl "http://vop.baidu.com/server_api?lan=zh&cuid=chuqingqing&token=24.774d214d120fb13f55b05b679004fab4.2592000.1427454902.282335-5415587" -H Content-Type:audio/wav;rate=8000 -H Expect: --data-bin @test.pcm -v