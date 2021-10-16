# Table Structure

1.user

create table user(

id

uuid

user_name

passwd

nick_name

);



2.tweet(

id

uuid

time

user_uuid

content_text

//content_image

//content_video

);



3.comment(

id

uuid

time

tweet_uuid

user_uuid

content_text

);



4.like(

id

uuid

time

like_uuid

user_uuid

);



5.retweet(

id

uuid

time

retweet_uuid

user_uuid

);



6.subscribe( //关注关系表

id

uuid

user_uuid

followed_uuid

)



/* 时间戳 redis

7.push_pool(

id

user_uuid //被推送者的uuid

content[] //uuid

)



8.pull_pool(

id

user_uuid //我自己的uuid，每个用户保有用于存放最新的内容

content[]

)



