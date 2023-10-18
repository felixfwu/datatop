WITH temp_space AS
 (SELECT a.tablespace_name, total_bytes, file_bytes, use_bytes
    FROM (SELECT tablespace_name,
                 SUM(decode(maxbytes, 0, bytes, maxbytes)) total_bytes
            FROM dba_temp_files
           GROUP BY tablespace_name) a,
         (SELECT tablespace_name, SUM(bytes) file_bytes
            FROM dba_temp_files
           GROUP BY tablespace_name) b,
         (SELECT tablespace_name, SUM(bytes_used) use_bytes
            FROM v$temp_extent_pool
           GROUP BY tablespace_name) c
   WHERE a.tablespace_name = b.tablespace_name
     AND b.tablespace_name = c.tablespace_name),
undo_space AS
 (SELECT a.tablespace_name, total_bytes, file_bytes, use_bytes
    FROM (SELECT tablespace_name,
                 SUM(decode(maxbytes, 0, bytes, maxbytes)) total_bytes
            FROM dba_data_files
           WHERE tablespace_name IN
                 (SELECT tablespace_name
                    FROM dba_tablespaces
                   WHERE contents = 'UNDO')
           GROUP BY tablespace_name) a,
         (SELECT tablespace_name, SUM(bytes) file_bytes
            FROM dba_data_files
           WHERE tablespace_name IN
                 (SELECT tablespace_name
                    FROM dba_tablespaces
                   WHERE contents = 'UNDO')
           GROUP BY tablespace_name) b,
         (SELECT tablespace_name, sum(bytes) use_bytes
            FROM dba_undo_extents
           WHERE status IN ('ACTIVE', 'UNEXPIRED')
           GROUP BY tablespace_name) c
   WHERE a.tablespace_name = b.tablespace_name
     AND b.tablespace_name = c.tablespace_name),
data_space AS
 (SELECT a.tablespace_name, total_bytes, file_bytes, use_bytes
    FROM (SELECT tablespace_name,
                 SUM(decode(maxbytes, 0, bytes, maxbytes)) total_bytes
            FROM dba_data_files
           WHERE tablespace_name IN
                 (SELECT tablespace_name
                    FROM dba_tablespaces
                   WHERE contents = 'PERMANENT')
           GROUP BY tablespace_name) a,
         (SELECT tablespace_name, SUM(bytes) file_bytes
            FROM dba_data_files
           WHERE tablespace_name IN
                 (SELECT tablespace_name
                    FROM dba_tablespaces
                   WHERE contents = 'PERMANENT')
           GROUP BY tablespace_name) b,
         (SELECT tablespace_name, sum(bytes) use_bytes
            FROM dba_segments
           GROUP BY tablespace_name) c
   WHERE a.tablespace_name = b.tablespace_name
     AND b.tablespace_name = c.tablespace_name)
SELECT tablespace_name,
       total_bytes,
       file_bytes,
       use_bytes / total_bytes * 100 use_pct
  FROM temp_space
UNION ALL
SELECT tablespace_name,
       total_bytes,
       file_bytes,
       use_bytes / total_bytes * 100 use_pct
  FROM undo_space
UNION ALL
SELECT tablespace_name,
       total_bytes,
       file_bytes,
       use_bytes / total_bytes * 100 use_pct
  FROM data_space
 ORDER BY use_pct desc;