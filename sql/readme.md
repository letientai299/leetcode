# Helpful SQL tricks

## `WITH RECURSIVE` CTE

Helpful to dynamically create a column of data to join with existing data.

```sql
WITH RECURSIVE cte (n) AS
(
  SELECT 1
  UNION ALL
  SELECT n + 1 FROM cte WHERE n < 5
)
SELECT * FROM cte;
```

Ref: https://dev.mysql.com/doc/refman/8.0/en/with.html#common-table-expressions-recursive

For example, when solving [1651 Hopper Company Queries III](https://leetcode.com/problems/hopper-company-queries-iii/),
we need to have a a list of `start_days` from `2020-01-01` to `2020-10-01` to
group the existed data in windows of 3 months. Below is my solution:

```sql
WITH RECURSIVE
    start_days (day) AS
        (
            SELECT '2020-01-01'
            UNION ALL
            SELECT date_add(day, interval 1 month)
            FROM start_days
            WHERE month(day) < 10
        ),
    res as (select month(sd.day)                    as month,
                   round(sum(ride_distance) / 3, 2) as average_ride_distance,
                   round(sum(ride_duration) / 3, 2) as average_ride_duration
            from start_days sd
                     left outer join Rides r
                                     on r.requested_at between
                                         sd.day
                                         and date_sub(date_add(sd.day, interval 3 month), interval 1 day)
                     join AcceptedRides ar on r.ride_id = ar.ride_id
            group by month(sd.day)
            order by month(sd.day)
    )
select month(sd.day)                        as month,
       ifnull(res.average_ride_distance, 0) as average_ride_distance,
       ifnull(res.average_ride_duration, 0) as average_ride_duration
from start_days sd
         left join res on month(sd.day) = res.month;
```
