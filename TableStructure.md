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

time_created

creator_uuid

is_retweet

content_text

origin

//content_image

//content_video

);



3.comment(

id

uuid

time_created

tweet_uuid

creator_uuid

content_text

);



4.like(

id

uuid

time_created

tweet_uuid

creator_uuid

);



5.retweet(

id

uuid

time

retweet_uuid

creator_uuid

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



