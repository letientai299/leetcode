-- https://leetcode.com/problems/find-the-subtasks-that-did-not-execute
-- Accepted solution.
WITH RECURSIVE cte (
    t,
    s
) AS (
    SELECT
        task_id,
        subtasks_count
    FROM
        Tasks
    UNION ALL
    SELECT
        t,
        s - 1
    FROM
        cte
    WHERE
        s > 1
)
SELECT
    t AS task_id,
    s AS subtask_id
FROM
    cte
    LEFT JOIN Executed e ON subtask_id = cte.s
        AND task_id = cte.t
WHERE
    e.subtask_id IS NULL;

-- I've almost solved this problem on my own with below query
--
--     with recursive cte(t, s) as (
--         select 1, 1
--         union all
--         select IF(s = (select subtasks_count from Tasks where task_id = t), t + 1, t),
--                IF(s = (select subtasks_count from Tasks where task_id = t), 1, s + 1)
--         from cte
--         where t < (select max(task_id) from Tasks)
--            or s < (select subtasks_count from Tasks where task_id = t)
--     )
--     select t as task_id, s as subtask_id
--     from cte
--     where not exists(select * from Executed where cte.t = task_id and cte.s = subtask_id);
--
-- That query has 2 issues causing a runtime error. First, we start at (1,1),
-- which might not even exist. Thus, it hits @@cte_max_recursion_depth limit.
-- Even if the limit is higher, it still fails with wrong result.
