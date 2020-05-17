https://sqliteonline.com/  
https://dev.mysql.com/downloads/windows/installer/


Выбор объектов, которые включены в максимальное число групп
``` sql
SELECT goods_id AS id, name FROM tags_goods 
	INNER JOIN goods ON goods.id = tags_goods.goods_id
	GROUP BY goods_id
	HAVING COUNT(tag_id) = (
		SELECT COUNT(goods_id) AS cnt FROM tags_goods
		GROUP BY goods_id
		ORDER BY cnt DESC
        LIMIT 1
	)
```
