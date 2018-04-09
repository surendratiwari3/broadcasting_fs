# broadcasting_fs

## To Do
1. All configuration will go into toml
2. From redis we will read the caller and callee both
3. Retry will be set based on configuration
4. Playback will be played using originate only
5. Getting the gateway list using distributor
6. Creating Campaign and then read the following thing
   1. start time
   2. end time
   3. Repeat (every day)
   4. wav file
   5. cps
   6. forced-callerid
   8. csv that will contain the caller and callee
   9. insert all this into redis
   10. concurrent call limit
7. API that will stop/start/restart the campaign
8. Successfull and failed one to redis
9. Multiple FreeSWITCH
10. redis cluster

Broadcast-Campaign
1. start-time (for example we need to run it daily morning 10 to evening 4)
2. caller-id (did-number) [comma seprated value] so we will take them one by one random
3. playback-file (currently it will be wav with sample rate 8khz)
4. call-duration
5. stop-time (for example we need to run it daily morning 10 to evening 4)
6. repeat_days [monday,tuesday,wednesday]
7. cps [default will be 2]

Upload CSV:
1. select compaign  
2. select csv file to upload
3. upload button
4. on upload button it will upload the csv file to server 

Redis Push Script:
1. This script will push the data to redis 

Redis:
1. LPUSH compaign_name phone_number

Multiple FreeSWITCH Server
1. fs1 (distributor) <------ go_esl (manual gateway creation) 




   
