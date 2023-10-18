package room

var GetFreeRooms = `
$needed_time = ?;
$default_time_from = CAST('2020-01-01T06:00:00.000000+3:00' AS TIMESTAMP);
$default_time_to = CAST('2020-01-01T23:00:00.000000+3:00' AS TIMESTAMP);

$r1 = (SELECT *
FROM rooms
WHERE rooms.is_available = true
  AND rooms.is_laboratory = false
  AND rooms.building = ?);

$r2 = (SELECT lessons.*
FROM lessons
         JOIN $r1 as r1
ON lessons.room_id = r1.id
WHERE lessons.week = ? OR lessons.week = 0 AND lessons.week_day = ?);

$r3 = (SELECT r1.*
FROM
    $r1 AS r1
    LEFT JOIN
    (SELECT * FROM $r2 as lessons WHERE
    lessons.time_from <= $needed_time
    AND lessons.time_to >= $needed_time
    ) as lessons
ON r1.id = lessons.room_id
WHERE lessons.id IS NULL);

$r4 = (SELECT r3.id         as id,
       r3.name              as name,
       r3.building          as building,
       MAX(lessons.time_to) as time_from
FROM $r3 AS r3 LEFT JOIN (SELECT * FROM $r2 as lessons WHERE lessons.time_to < $needed_time) as lessons
ON r3.id = lessons.room_id
GROUP BY r3.id, r3.name, r3.building);

$r5 = (SELECT r4.id                  as id,
       r4.name                as name,
       r4.building            as building,
       r4.time_from           as time_from,
       MIN(lessons.time_from) as time_to
FROM $r4 AS r4 LEFT JOIN (SELECT * FROM $r2 as lessons WHERE lessons.time_from > $needed_time) as lessons
ON r4.id = lessons.room_id
GROUP BY r4.id, r4.name, r4.building, r4.time_from);

SELECT id,
       name,
       building,
       COALESCE(time_from, $default_time_from) as time_from,
       COALESCE(time_to, $default_time_to)     as time_to,
       FROM $r5 as r5;
`
