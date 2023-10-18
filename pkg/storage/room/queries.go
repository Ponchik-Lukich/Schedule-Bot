package room

var GetFreeRooms = `
WITH consts AS (SELECT CAST(? AS TIMESTAMP) AS needed_time,
                       CAST('2020-01-01T03:00:00.000000Z' AS TIMESTAMP) AS default_time_from,
                       CAST('2020-01-01T20:00:00.000000Z' AS TIMESTAMP) AS default_time_to),
     r1 AS (SELECT *
            FROM rooms
            WHERE is_available = true
               AND building NOT IN (
                  'Корпус Спортивный', 
                  'Корпус Нет корпуса', 
                  'Корпус ФМ', 
                  'Корпус ХХ', 
                  'Корпус 44а', 
                  'Корпус 6А'
               )
               AND name NOT LIKE 'каф%'
               AND building = ?
               AND is_laboratory = false),
     r2 AS (SELECT lessons.*
            FROM lessons
                     JOIN r1
                          ON lessons.room_id = r1.id
            WHERE (lessons.week = ? OR lessons.week = 0)
              AND lessons.week_day = ?),
     r3 AS (SELECT r1.*
            FROM r1
                     CROSS JOIN consts
                     LEFT JOIN r2 AS lessons ON r1.id = lessons.room_id
                AND lessons.time_from <= consts.needed_time
                AND lessons.time_to >= consts.needed_time
            WHERE lessons.id IS NULL),
     r4 AS (SELECT r3.id, r3.name, r3.building, MAX(lessons.time_to) AS time_from
            FROM r3
                     CROSS JOIN consts
                     LEFT JOIN r2 AS lessons ON r3.id = lessons.room_id
                AND lessons.time_to < consts.needed_time
            GROUP BY r3.id, r3.name, r3.building),
     r5 AS (SELECT r4.id, r4.name, r4.building, r4.time_from, MIN(lessons.time_from) AS time_to
            FROM r4
                     CROSS JOIN consts
                     LEFT JOIN r2 AS lessons ON r4.id = lessons.room_id
                AND lessons.time_from > consts.needed_time
            GROUP BY r4.id, r4.name, r4.building, r4.time_from)

SELECT r5.id,
       r5.name,
       r5.building,
       COALESCE(r5.time_from, consts.default_time_from) AS time_from,
       COALESCE(r5.time_to, consts.default_time_to)     AS time_to
FROM r5,
     consts;
`
