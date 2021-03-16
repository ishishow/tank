use dojo_api;

SET NAMES utf8mb4;

INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("1","ishishow-auth-token1","ishishow",10,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("2","ishishow-auth-token2","いししょう",20,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("3","ishishow-auth-token3","nishio",30,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("4","ishishow-auth-token4","240",30,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("5","ishishow-auth-token5","fukuda",40,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("6","ishishow-auth-token6","ariga",40,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("7","ishishow-auth-token7","henmi",40,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("8","ishishow-auth-token8","mentor",50,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("9","ishishow-auth-token9","ww",50,0);
INSERT INTO `user` (`id`,`auth_token`,`name`,`high_score`,`coin`) VALUES ("10","ishishow-auth-token10","ishishow",50,0);

INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("1","インスタント戦車",1);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("2","山吹色の戦車R",1);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("3","緋色の戦車SR",2);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("4","紺碧の戦車SR",2);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("5","黄金の戦車SSR",3);

INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("6","インスタント武器",1);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("7","山吹色の武器SR",1);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("8","緋色の武器SR",2);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("9","紺碧の武器SR",2);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("10","黄金の武器SSR",3);
INSERT INTO `item` (`id`,`name`,`rarity`) VALUES ("11","レールガンSSR",3);

INSERT INTO `skin` (`id`,`item_id`,`hit_point`,`speed`) VALUES ("s_green","1",100,15.0);
INSERT INTO `skin` (`id`,`item_id`,`hit_point`,`speed`) VALUES ("s_yellow","2",120,16.0);
INSERT INTO `skin` (`id`,`item_id`,`hit_point`,`speed`) VALUES ("s_red","3",160,16.0);
INSERT INTO `skin` (`id`,`item_id`,`hit_point`,`speed`) VALUES ("s_blue","4",140,18.0);
INSERT INTO `skin` (`id`,`item_id`,`hit_point`,`speed`) VALUES ("s_gold","5",200,18.0);

INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_green","6",10,50,5.0,1500.0);
INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_yellow","7",10,75,5.0,1500.0);
INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_red","8",10,100,5.0,1500.0);
INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_blue","9",10,100,5.0,1500.0);
INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_gold","10",10,150,5.0,1500.0);
INSERT INTO `weapon` (`id`,`item_id`,`ballet`,`attack`,`reload`,`speed`) VALUES ("w_railgun","11",10,250,5.0,1000.0);


INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("1",0);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("2",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("3",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("4",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("5",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("6",0);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("7",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("8",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("9",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("10",3);
INSERT INTO `gacha_probability` (`item_id`,`ratio`) VALUES ("11",3);