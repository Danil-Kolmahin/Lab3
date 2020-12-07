\i G:/3_cource/'Software _Architecture'/Lab3/db/Balancers.sql
\i G:/3_cource/'Software _Architecture'/Lab3/db/Machines.sql
\i G:/3_cource/'Software _Architecture'/Lab3/db/ConnectToBalancers.sql

SELECT * FROM Machines LIMIT 10;

SELECT balancer_id, COUNT(*) FROM ConnectToBalancers GROUP BY balancer_id LIMIT 10;
SELECT machine_id FROM ConnectToBalancers WHERE balancer_id = 71 LIMIT 10;

SELECT name, isUsed FROM ConnectToBalancers
INNER JOIN machines m on m.id = connecttobalancers.machine_id
WHERE machine_id = 72 LIMIT 10;



UPDATE Machines SET isUsed = false WHERE id = 7;

SELECT balancer_id AS "id",
       array_agg(machine_id)::varchar(255) AS "usedMachenes",
       COUNT(*) AS "totalMachenes"
FROM ConnectToBalancers, Machines
WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = true
GROUP BY balancer_id;



