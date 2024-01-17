# leinvetrorie
Nice little inventory system for all your inventory needs.

![Start page image](https://github.com/Seth-Buchanan/leinvetrorie/blob/2145383ffb1b569fc1424eb5d06c0bf5e3f664be/start.png?raw=true)

## Add and remove item quantities 
Wow! a cool new sleek minimal useful terminal user interface that that adds *and* removes items!

## None of that administration bloatware!
Our users know how to use sqlite3 to add and remove items from the table the program makes.

### Adding Items
```SQL
INSERT INTO 
ourtable
(name, quantity, replacement_level) 
VALUES
("Lemons",80, 32),
("Orangutans",120, 83),
("Bollards",25, 5),
("Turtles",25, 73),
("Basketballs",25, 4);
```

### Listing Items
```SQL
SELECT * FROM ourtable;
```
output:
```
1|Lemons|80|32
2|Orangutans|120|83
3|Bollards|25|5
4|Turtles|25|73
5|Basketballs|25|4
```

### Removing Items
```SQL
DELETE FROM ourtable WHERE id = 1;
```

# Build steps

``` bash
$ git clone https://github.com/Seth-Buchanan/leinvetrorie.git

$ cd leinvetrorie

$ go build .
```

# Usage
``` bash 
./main
```

* Select add or remove action

* Enter bin number (Default is 1-3)

* enter the amount you would like to add or remove

and most importantly, enjoy.
