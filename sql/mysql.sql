-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: route-planner
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

create database route_planner;
use route_planner;
--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `cid` bigint NOT NULL,
  `name` text NOT NULL,
  `sid` int NOT NULL,
  `content` text,
  `likes` int DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1622175807059267584,'np123',1,'wevewv',2,'2023-04-08 00:23:24'),(1622175835354042368,'admin',1,'everver',2,'2023-04-08 00:23:24'),(1622175837895790592,'np123',2,'erverv',2,'2023-04-08 00:23:24'),(1622175839275716608,'admin',2,'erverber',2,'2023-04-08 00:23:24'),(1622175840655642624,'np123',3,'berb',2,'2023-04-08 00:23:24'),(1622175843977531392,'admin',3,'erberberbe',2,'2023-04-08 00:23:24'),(1622176229627006976,'admin',3,'rberb',2,'2023-04-08 00:23:24'),(1622954288307245056,'admin',1,'23t234t',0,'0000-00-00 00:00:00'),(1622954403289894912,'admin',1,'评论测试！！！',0,'0000-00-00 00:00:00'),(1622954867771314176,'NP123',1,'评论！',0,'2023-02-07 21:46:08'),(1622954952584335360,'NP123',2,'好好哈',0,'2023-02-07 21:46:28'),(1622955030929739776,'NP123',2,'刷新测试',0,'2023-02-07 21:46:47'),(1622955062928084992,'NP123',3,'加油\n！',0,'2023-02-07 21:46:55'),(1623211719859900416,'测试用户',1,'很好!',0,'2023-02-08 14:46:46'),(1623211764080447488,'测试用户',2,'哈哈哈哈哈',0,'2023-02-08 14:46:57'),(1623227025835692032,'NP123',1,'qefg4g54j65j',0,'2023-02-08 15:47:36'),(1623227613709340672,'admin',1,'达瓦达瓦达瓦',0,'2023-02-08 15:49:56'),(1623227701043138560,'admin',2,'很好很好',0,'2023-02-08 15:50:17'),(1623245358513524736,'admin',2,'傻逼',0,'2023-02-08 17:00:27'),(1623589122163609600,'test',1,'segrshtdn',0,'2023-02-09 15:46:26'),(1623591290836881408,'npneo',1,'啊啊啊啊',0,'2023-02-09 15:55:03');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `price`
--

DROP TABLE IF EXISTS `price`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `price` (
  `sid` int DEFAULT NULL,
  `aimed` text,
  `price` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `price`
--

LOCK TABLES `price` WRITE;
/*!40000 ALTER TABLE `price` DISABLE KEYS */;
INSERT INTO `price` VALUES (1,'学生',4.99),(1,'成人',5.8),(2,'学生',10),(2,'成人',5),(2,'老人',2),(3,'学生',20),(3,'老师',30),(3,'教授',15),(3,'脑瘫 ',2);
/*!40000 ALTER TABLE `price` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sites`
--

DROP TABLE IF EXISTS `sites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sites` (
  `sid` int NOT NULL,
  `sname` text,
  `description` text,
  `pic` text,
  `website` text,
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sites`
--

LOCK TABLES `sites` WRITE;
/*!40000 ALTER TABLE `sites` DISABLE KEYS */;
INSERT INTO `sites` VALUES (1,'华东理工大学','华东理工大学（East China University of Science and Technology），简称华理（ECUST），坐落于上海市，是中华人民共和国教育部直属的一所学科设置涵盖11个学科门类的全国重点大学，位列国家“双一流”、“211工程”、“985工程优势学科创新平台”，入选国家建设高水平大学公派研究生项目、高等学校创新能力提升计划、高等学校学科创新引智计划、基础学科拔尖学生培养计划2.0、双万计划、卓越工程师教育培养计划、新工科研究与实践项目、国家大学生创新性实验计划、国家级大学生创新创业训练计划、全国深化创新创业教育改革示范高校、中国政府奖学金来华留学生接收院校，为全国16所工科重点大学科技工作研讨会、高水平行业特色大学优质资源共享联盟成员。 [1-2]   [182] \n华东理工大学原名华东化工学院，1952年由交通大学、震旦大学、大同大学、东吴大学、江南大学等校化工系组建而成，1956年被定为全国首批招收研究生的学校之一；1960年被确定为教育部直属的全国重点大学，1993年改名华东理工大学；1996年成为国家“211工程”重点建设高校。 [1] \n截至2023年1月，学校有徐汇、奉贤和金山三个校区，占地面积2532亩，各类建筑总面积95万平方米，馆藏图书353万册；设有17个学院，开设本科招生专业72个；有博士后科研流动站14个、一级学科博士学位授权点18个、博士专业学位授权点5个、一级学科硕士学位授权点31个、硕士专业学位授权点17个；有在校全日制学生2.9万余人，其中，本科生16764人，教职员工约3000余人。','http://localhost:8080/images/1.jpg','ecust.edu.cn'),(2,'外滩','外滩（英文：The Bund；上海话拼音：nga thae），位于上海市黄浦区的黄浦江畔，即外黄浦滩，为中国历史文化街区。1844年（清道光廿四年）起，外滩这一带被划为英国租界，成为上海十里洋场的真实写照，也是旧上海租界区以及整个上海近代城市开始的起点。\n外滩全长1.5千米，南起延安东路，北至苏州河上的外白渡桥，东面即黄浦江，西面是旧上海金融、外贸机构的集中地。上海辟为商埠以后，外国的银行、商行、总会、报社开始在此云集，外滩成为全国乃至远东的金融中心。民国三十二年（1943年）8月，外滩随交还上海公共租界于汪伪国民政府，结束长达百年的租界时期，于民国三十四年（1945年）拥有正式路名中山东一路。\n外滩矗立着52幢风格迥异的古典复兴大楼，素有外滩万国建筑博览群之称，是中国近现代重要史迹及代表性建筑，上海的地标之一，1996年11月，中华人民共和国国务院将其列入第四批全国重点文物保护单位。\n2018年3月，上海外滩在全面推进“第一立面”（即临江建筑群）功能置换的基础上，同步启动了“第二立面”（即非临江的外滩建筑群）功能置换工作。','http://localhost:8080/images/2.jpg','awdcawdtwdawdaw'),(3,'np123','## 例子\n\n> **I can think of no better illustration of the view than the case of our parents.**\n\n口语：Personally speaking, I prefer XXX, just because XXX saves me a lot of time/is a waste of time. \n\n**Time is primary productive force.**\n\n* **论据-节省时间：As a senior majoring in computer science, I have a packed schedule with daily tasks such as participating in algorithm competitions, studying new programming languages, and preparing for tomorrow\'s classes. I spend most of my day in labs and the library and also have to fit in internships in my spare time. Thus, I have no time to waste.**\n\n* **论据-减轻压力：It can alleviate our stress and maintain both our physical and emotional well-being. The pressure confronted by young people is chokingly high, leading them to feel exhausted and stressed out,especially ...**\n\n  \n\n## 重要性\n\n* **It seems almost inevitable** that the historical information **is pivotal for anybody in anytime.**\n* **It is imperative for all** school teachers **to** renew their own knowledge by taking courses periodically.\n\n## 论证\n\n* **A>B(好处): The advantage derived from the voluminous Internet world far outweigh those benefits that** learner gained from the knowledge of what has been printed in the textbook.\n* **A>B(重要性):Whereas nowadays, the prevalence of the Internet has given rise to the spread of** history content**, which considerably undermines and eliminate the importance of the history course**\n* **B不好:Fascination with something often comes at a cost.** If we fail to manage our time and let our mania encroach on the time to work, ..., **undesirable consequences would occur.**\n* **从基础来说：Let us get down to fundamentals and agree that ...** \n\n## 效率\n\n* **Implementing** this strategy **will result in significant enhancements in productivity and efficiency.**\n\n## 价值观\n\n* **can provide students with an enlightening experience that expands their perspective and enhances their understanding of the world.**\n\n## 其他\n\n* Interests like the light in the dark, **it will stimulate us to keep going** in the long run.\n\n* **The rise of the population is swift**\n* **invade our thoughts or interrupt** the pleasure of meeting friends for a quiet chat\n\n* **On the other hand**, the study of world history **has contributed greatly to** the mutual understanding and cooperation of world nations.\n* **It is conceivable that ...**\n* **It is not uncommon that**\n* **physical exercise are conducive to people\'s physical condition,** which intensifies their resistance against fatigue.\n* For me, studying computer **have the same peculiar fascination that** physics for those famous people like Einstein and Newton.\n\n','http://localhost:8080/images/3.jpg','github.com');
/*!40000 ALTER TABLE `sites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sites_tags`
--

DROP TABLE IF EXISTS `sites_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sites_tags` (
  `sid` int DEFAULT NULL,
  `tagid` int DEFAULT NULL,
  `level` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sites_tags`
--

LOCK TABLES `sites_tags` WRITE;
/*!40000 ALTER TABLE `sites_tags` DISABLE KEYS */;
INSERT INTO `sites_tags` VALUES (1,1,0.5),(1,2,0.2),(2,3,0.7),(3,1,0.8),(2,2,1),(3,3,1.2);
/*!40000 ALTER TABLE `sites_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags` (
  `tagid` int DEFAULT NULL,
  `name` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'学校'),(2,'景点'),(3,'公园'),(4,'小区');
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `name` text,
  `password` text,
  `email` text,
  `avatar` text,
  `createtime` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('admin','123456','np123greatest@gmail.com','http://localhost:8080/images/1.jpg','2023-02-06 20:39:39'),('NP123','Abc12345678','38424@sina.com','http://localhost:8080/images/3.jpg','2023-02-06 20:54:09'),('测试用户','123456','38424@sina.com','http://localhost:8080/images/1675838784.jpg','2023-02-08 14:46:24'),('test','123456','38424@sina.com','http://localhost:8080/images/1675928760.jpg','2023-02-09 15:46:00'),('admin2','123456','38424@sina.com','http://localhost:8080/images/1675929224.jpg','2023-02-09 15:53:45'),('npneo','Abc1234567890','38424@sina.com','http://localhost:8080/images/1675929284.jpg','2023-02-09 15:54:45');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-02-09 22:13:52
