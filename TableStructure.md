# Table Structure

1.user

create table user(

id

uuid

user_name

passwd

nickname

);



2.tweet(

id

uuid

time

user_uuid

content

);



3.comment(

id

uuid

time

tweet_uuid

user_uuid

content_text

//content_image

//content_video

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



6.subscribe