https://sqliteonline.com/  
https://dev.mysql.com/downloads/windows/installer/


Выбор объектов, которые включены в максимальное число групп
``` sql
SELECT goods_id as id, name FROM tags_goods 
	inner join goods on goods.id = tags_goods.goods_id
	GROUP BY goods_id
	having COUNT(tag_id) = (
		select count(goods_id) as cnt from tags_goods
		group by goods_id
		order by cnt desc
        limit 1
	)
```
