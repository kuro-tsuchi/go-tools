create table if not exists test.order
(
  id String,
  user_id Int64,
  money Decimal32(2) ,
  vip String
 
) engine = MergeTree
order by id;
