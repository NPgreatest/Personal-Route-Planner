-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: route_planner
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
INSERT INTO `comment` VALUES (1622175807059267584,'np123',1,'wevewv',2,'2023-04-08 00:23:24'),(1622175835354042368,'admin',1,'everver',2,'2023-04-08 00:23:24'),(1622175837895790592,'np123',2,'erverv',2,'2023-04-08 00:23:24'),(1622175839275716608,'admin',2,'erverber',2,'2023-04-08 00:23:24'),(1622175840655642624,'np123',3,'berb',2,'2023-04-08 00:23:24'),(1622175843977531392,'admin',3,'erberberbe',2,'2023-04-08 00:23:24'),(1622176229627006976,'admin',3,'rberb',2,'2023-04-08 00:23:24'),(1622954288307245056,'admin',1,'23t234t',0,'0000-00-00 00:00:00'),(1622954403289894912,'admin',1,'评论测试！！！',0,'0000-00-00 00:00:00'),(1622954867771314176,'NP123',1,'评论！',0,'2023-02-07 21:46:08'),(1622954952584335360,'NP123',2,'好好哈',0,'2023-02-07 21:46:28'),(1622955030929739776,'NP123',2,'刷新测试',0,'2023-02-07 21:46:47'),(1622955062928084992,'NP123',3,'加油\n！',0,'2023-02-07 21:46:55'),(1623211719859900416,'测试用户',1,'很好!',0,'2023-02-08 14:46:46'),(1623211764080447488,'测试用户',2,'哈哈哈哈哈',0,'2023-02-08 14:46:57'),(1623227025835692032,'NP123',1,'qefg4g54j65j',0,'2023-02-08 15:47:36'),(1623227613709340672,'admin',1,'达瓦达瓦达瓦',0,'2023-02-08 15:49:56'),(1623227701043138560,'admin',2,'很好很好',0,'2023-02-08 15:50:17'),(1623245358513524736,'admin',2,'傻逼',0,'2023-02-08 17:00:27'),(1623589122163609600,'test',1,'segrshtdn',0,'2023-02-09 15:46:26'),(1623591290836881408,'npneo',1,'啊啊啊啊',0,'2023-02-09 15:55:03'),(1624679597205360640,'zzz',1,'dawdwada',0,'2023-02-12 15:59:36');
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
INSERT INTO `price` VALUES (1,'学生',4.99),(1,'成人',5.8),(2,'学生',10),(2,'成人',5),(2,'老人',2),(3,'学生',20),(3,'老师',30),(3,'教授',15),(3,'脑瘫 ',2),(4,'成人票',45),(4,'学生票',22),(4,'老年票',37),(6,'成人大门票',15),(6,'成人联票(含大门票及售票专类园)',40),(6,'儿童大门票',7.5),(6,'童联票(含大门票及售票专类园)',20),(7,'巅峰票：成人票',220),(7,'巅峰票：儿童票',110),(7,'速通票：成人票',299),(7,'速通票：儿童票',199),(7,'超值票：成人票',199),(7,'超值票：儿童票',99),(7,'至尊VIP套票',399),(8,'外滩轮渡',2),(8,'外滩观光隧道单程票成人票',45);
/*!40000 ALTER TABLE `price` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rating`
--

DROP TABLE IF EXISTS `rating`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rating` (
  `name` text,
  `rate` json DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rating`
--

LOCK TABLES `rating` WRITE;
/*!40000 ALTER TABLE `rating` DISABLE KEYS */;
INSERT INTO `rating` VALUES ('admin','{\"Name\": \"admin\", \"rates\": [{\"sid\": 1, \"rate\": 2}, {\"sid\": 2, \"rate\": 3}]}'),('test','{\"Name\": \"test\", \"rates\": [{\"sid\": 1, \"rate\": 2}, {\"sid\": 2, \"rate\": 3}]}'),('NP123','{\"name\": \"NP123\", \"rates\": [{\"sid\": 1, \"rate\": 4}, {\"sid\": 2, \"rate\": 17}]}');
/*!40000 ALTER TABLE `rating` ENABLE KEYS */;
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
INSERT INTO `sites` VALUES (1,'华东理工大学','华东理工大学（East China University of Science and Technology），简称华理（ECUST），坐落于上海市，是中华人民共和国教育部直属的一所学科设置涵盖11个学科门类的全国重点大学，位列国家“双一流”、“211工程”、“985工程优势学科创新平台”，入选国家建设高水平大学公派研究生项目、高等学校创新能力提升计划、高等学校学科创新引智计划、基础学科拔尖学生培养计划2.0、双万计划、卓越工程师教育培养计划、新工科研究与实践项目、国家大学生创新性实验计划、国家级大学生创新创业训练计划、全国深化创新创业教育改革示范高校、中国政府奖学金来华留学生接收院校，为全国16所工科重点大学科技工作研讨会、高水平行业特色大学优质资源共享联盟成员。 [1-2]   [182] \n华东理工大学原名华东化工学院，1952年由交通大学、震旦大学、大同大学、东吴大学、江南大学等校化工系组建而成，1956年被定为全国首批招收研究生的学校之一；1960年被确定为教育部直属的全国重点大学，1993年改名华东理工大学；1996年成为国家“211工程”重点建设高校。 [1] \n截至2023年1月，学校有徐汇、奉贤和金山三个校区，占地面积2532亩，各类建筑总面积95万平方米，馆藏图书353万册；设有17个学院，开设本科招生专业72个；有博士后科研流动站14个、一级学科博士学位授权点18个、博士专业学位授权点5个、一级学科硕士学位授权点31个、硕士专业学位授权点17个；有在校全日制学生2.9万余人，其中，本科生16764人，教职员工约3000余人。','http://20.89.130.136:8080/images/1.jpg','http://ecust.edu.cn'),(2,'外滩','外滩（英文：The Bund；上海话拼音：nga thae），位于上海市黄浦区的黄浦江畔，即外黄浦滩，为中国历史文化街区。1844年（清道光廿四年）起，外滩这一带被划为英国租界，成为上海十里洋场的真实写照，也是旧上海租界区以及整个上海近代城市开始的起点。\n外滩全长1.5千米，南起延安东路，北至苏州河上的外白渡桥，东面即黄浦江，西面是旧上海金融、外贸机构的集中地。上海辟为商埠以后，外国的银行、商行、总会、报社开始在此云集，外滩成为全国乃至远东的金融中心。民国三十二年（1943年）8月，外滩随交还上海公共租界于汪伪国民政府，结束长达百年的租界时期，于民国三十四年（1945年）拥有正式路名中山东一路。\n外滩矗立着52幢风格迥异的古典复兴大楼，素有外滩万国建筑博览群之称，是中国近现代重要史迹及代表性建筑，上海的地标之一，1996年11月，中华人民共和国国务院将其列入第四批全国重点文物保护单位。\n2018年3月，上海外滩在全面推进“第一立面”（即临江建筑群）功能置换的基础上，同步启动了“第二立面”（即非临江的外滩建筑群）功能置换工作。','http://20.89.130.136:8080/images/2.jpg','http://baidu.com'),(3,'np123','## 例子\n\n> **I can think of no better illustration of the view than the case of our parents.**\n\n口语：Personally speaking, I prefer XXX, just because XXX saves me a lot of time/is a waste of time. \n\n**Time is primary productive force.**\n\n* **论据-节省时间：As a senior majoring in computer science, I have a packed schedule with daily tasks such as participating in algorithm competitions, studying new programming languages, and preparing for tomorrow\'s classes. I spend most of my day in labs and the library and also have to fit in internships in my spare time. Thus, I have no time to waste.**\n\n* **论据-减轻压力：It can alleviate our stress and maintain both our physical and emotional well-being. The pressure confronted by young people is chokingly high, leading them to feel exhausted and stressed out,especially ...**\n\n  \n\n## 重要性\n\n* **It seems almost inevitable** that the historical information **is pivotal for anybody in anytime.**\n* **It is imperative for all** school teachers **to** renew their own knowledge by taking courses periodically.\n\n## 论证\n\n* **A>B(好处): The advantage derived from the voluminous Internet world far outweigh those benefits that** learner gained from the knowledge of what has been printed in the textbook.\n* **A>B(重要性):Whereas nowadays, the prevalence of the Internet has given rise to the spread of** history content**, which considerably undermines and eliminate the importance of the history course**\n* **B不好:Fascination with something often comes at a cost.** If we fail to manage our time and let our mania encroach on the time to work, ..., **undesirable consequences would occur.**\n* **从基础来说：Let us get down to fundamentals and agree that ...** \n\n## 效率\n\n* **Implementing** this strategy **will result in significant enhancements in productivity and efficiency.**\n\n## 价值观\n\n* **can provide students with an enlightening experience that expands their perspective and enhances their understanding of the world.**\n\n## 其他\n\n* Interests like the light in the dark, **it will stimulate us to keep going** in the long run.\n\n* **The rise of the population is swift**\n* **invade our thoughts or interrupt** the pleasure of meeting friends for a quiet chat\n\n* **On the other hand**, the study of world history **has contributed greatly to** the mutual understanding and cooperation of world nations.\n* **It is conceivable that ...**\n* **It is not uncommon that**\n* **physical exercise are conducive to people\'s physical condition,** which intensifies their resistance against fatigue.\n* For me, studying computer **have the same peculiar fascination that** physics for those famous people like Einstein and Newton.\n\n','http://20.89.130.136:8080/images/3.jpg','github.com'),(4,'上海科技馆','<div>位于中国上海市浦东新区世纪大道2000号，花木行政文化中心区，世纪广场西侧，南邻世纪公园。国家5A级 景区。主馆占地面积6.8万多平方米，建筑面积9.8万平方米，分为11个风格各异的主题展区、4个高科技特种影院、3个古今中外科学家及其足迹的艺术长廊、2个主题特展和若干个临时展厅，它们共同为四方游客生动地演绎着\"自然、人、科技\"的永恒话题。</div></div>','http://20.89.130.136:8080/images/上海科技馆.jpg','http://www.quanjingke.com/jq/m/411'),(5,'城隍庙','<div>位于上海市黄浦区方浜中路，地铁十、十四号线途经。国家4A级旅游景区，为“长江三大庙”之一 。</div></div><div>城隍庙殿堂建筑属南方大式建筑，红墙泥瓦，庙内主体建筑由庙前广场、大殿、元辰殿，财神殿、慈航殿、城隍殿、娘娘殿组成。现门前存有1535年所建的牌坊，戏台为1865年建。</div>','http://20.89.130.136:8080/images/城隍庙.jpg','http://www.quanjingke.com/jq/m/67'),(6,'上海植物园','<div>位于徐汇区。国家4A级景区。园区分为展览区和绿化种植区，展览区建有植物进化区、环境保护区、人工生 态区、绿化示范区4个展出区和黄道婆庙游览区，各区下分以专类植物为主景的若干小区，配以园林建筑小品。园区主要收集引种长江中 下游野生植物，共收集活植物3500余种，6000多个品种。</div></div>','http://20.89.130.136:8080/images/上海植物园.jpg','http://www.quanjingke.com/jq/m/416'),(7,'东方明珠','<div>位于上海市浦东新区陆家嘴世纪大道1号，地处黄浦江畔，地铁二、十、十四号线途经。国家5A级景区，背拥陆家嘴地区现代化建筑楼群，与隔江的外滩万国建筑博览群交相辉映，是集都市观光、时尚餐饮、购物娱乐、历史陈列、浦江游览、会展演出、广播电视发射等多功能于一体的上海市标志性建筑之一。</div></div>','http://20.89.130.136:8080/images/东方明珠.jpg','http://www.quanjingke.com/jq/m/103'),(8,'外滩','<div>上海外滩，北起外白渡桥，南至延安东路，全程大约1.5公里，东边毗邻黄浦江，西侧与浦东金融中心陆家嘴隔江而对，是感受中国历史沧海桑田之变化的都市游览胜地。黄浦江从城市中心蜿蜒流过，东侧是有“万国博览建筑群”之称的旧上海建筑，是旧上海城市历史的见证；西侧为现代上海的地标建筑：金茂大厦、东方明珠、上海环球金融中心、上海中心大厦等，是中国现代化建设与改革开放成果的缩影。漫步外滩，我们可以感受到上海这座国际化大都市的历史与发展，更能体会到祖国在世界东方崛起带来的风云变幻。</div></div><div>白天游览外滩可以从外白渡桥出发，一路游览万国博览建筑群。外白渡桥本身也是著名的景点。这是一座简支钢桁架桥，同时也是一座全钢结构铆接桥梁和不等高桁架结构桥梁，有两孔。外白渡桥上部结构为下承式简支铆接钢珩架，下部木桩基础钢筋混凝土桥台和混凝土空心薄板桥墩桥支撑着上部下承式筒支铆接钢桁架，下部结构为木桩基础钢筋混凝土桥台和混凝土空心薄板桥墩。外白渡桥同时也是中国的第一座全钢结构铆接桥梁和仅存的不等高桁架结构桥梁自19世纪40年代租界被英法等国抢占后，外滩便成了一个主权区，西方列强以他们的方式经营、管理、建设租界，当商行、金融企业在外滩占有一席之地后，即大兴土木，营建公司大楼，而外滩的建筑大多也经过三次或三次以上的重建。20世纪，由于建筑技术的发展和经济实力的增长，外滩出现了多层和高层建筑，式样五花八门，诸如英国古典式、英国新古典式、英国文艺复兴式亚细亚大楼（原上海冶金设计院）、上海总会（今东风饭店）、浦发银行大楼（原汇丰银行大楼）、恰和大楼（今外贸局大楼）等，还有法国古典式、法国大住宅式、哥特式、巴洛克式、近代西方式、东印度式、折中主义式、中西掺合式等，呈现世界各国建筑共存的局面。因而，北起苏州河外白渡桥，南至中山东一路金陵东路的这一片建筑群，被誉为“万国建筑博览”。这些古典主义与现代主义并存的建筑，已成为了上海的象征。</div>','http://20.89.130.136:8080/images/外滩.jpg','http://www.quanjingke.com/jq/m/508'),(9,'中共一大纪念馆','<div>中共一大会址地处上海原法租界望志路106号（今兴业路76号），属典型上海石库门风格建筑，建于1920年秋。1921年7月23日，中国共产党第一次全国代表大会在此召开。中国共产党第一次全国代表大会通过了中国共产党的第一个纲领和第一个决议，选举产生了中央领导机构，宣告了中国共产党的诞生。1952年9月，中共一大会址修复，建立纪念馆并对外开放。1961年3月，国务院公布中共一大会址为全国重点文物保护单位。2017年10月31日，习近平总书记带领中央政治局常委集体瞻仰中共一大会址并在宣誓大厅庄严宣誓，回顾建党历史，重温入党誓词，宣示新一届党中央领导集体的坚定政治信念。</div></div><div>中共一大会址原楼于民国九年（1920年）夏秋间建，与左右紧邻4幢同类房屋同时建成，属贝勒路树德里（今黄陂南路374弄）一部分。为上海典型石库门式样建筑，外墙青红砖交错，镶嵌白色粉线，门楣有矾红色雕花，黑漆大门上配铜环，门框围以米黄色石条，门楣上部有拱形堆塑花饰。“中国共产党创建历史文物陈列”由三部分内容组成：中国党成立的历史背景；中国共产党早期组织的成立及其活动；中国共产党第一次全国代表大会召开的全过程。陈列室内还按照历史资料开辟了一个蜡像室，形象地刻画出当年出席中共“一大”会议的15位出席者（包括2位共产国际代表）围桌而坐、热烈讨论的生动场景。栩栩如生的蜡像人物增强了陈列的直观性和历史感染力，成为参观中的新热点。</div>','http://20.89.130.136:8080/images/中共一大纪念馆.jpg','https://360.zgyd1921.com/project/57/index1.html?content=0&&startscene=0&startactions=lookat(-291.47,88.21,149.4,0.91,0);'),(10,'陆家嘴','<div>地处上海市区中部，位于上海市浦东新区的黄浦江畔，地铁二号线、十四号线途经。两面环水，西面隔江与外滩 万国建筑博览群相望，北面隔江眺望北外滩。陆家嘴是上海国际金融中心的核心功能区，为多家跨国银行的中国（含港澳台）及东亚总部所在地。陆家嘴境内有东方明珠广播电视塔、上海中心大厦、上海环球金融中心、上海金茂大厦等现代化建筑群。</div></div>','http://20.89.130.136:8080/images/陆家嘴.jpg','http://www.quanjingke.com/jq/m/1279'),(11,'上海国际会议中心','<div>地处陆家嘴金融贸易中心，毗邻东方明珠电视塔，与外滩万国建筑群隔江相望，被评为建国五十年十 大经典建筑之一，素以举办大型国际会议、商务论坛而蜚声海内外。</div></div>','http://20.89.130.136:8080/images/上海国际会议中心.jpg','http://www.quanjingke.com/jq/m/409'),(12,'上海人民公园','<div>上海人民公园地处黄浦区，位于上海市中心，主园门位于南京西路231号。园内有南极石、张思德雕像和五卅纪念碑成为青少年爱国主义教育的标志，还设有艺术馆、迷你但齐全的\"欢乐谷\"，北面有国际饭店、大光明电影院等古董老建筑，南面隔着人民广场是摩天楼组成的城市新兴天际线。</div></div>','http://20.89.130.136:8080/images/上海人民公园.jpg','http://www.quanjingke.com/jq/m/413'),(13,'上海世纪广场','<div>位于浦东新区世纪大道南道，是上海最大的露天广场，也是唯一一个以时间为主题的雕塑景观广场。广场 整体建筑主要体现\"世纪之光\"这一主题。广场的入口是以日晷为原形设计的大型景观雕塑\"东方之光\"，突出跨世纪的时间主题，是雕塑艺术语言与现代高科技建筑语言的完美结合。</div></div>','http://20.89.130.136:8080/images/上海世纪广场.jpg','http://www.quanjingke.com/jq/m/411'),(14,'上海中山公园','<div>位于长宁区中心，地铁二、三、四号线途经，附近有来福士、龙之梦等大商场。1941年为纪念孙中山先生 而得名。公园以英国式自然造园风格为主，融中国园林艺术之精华，中西合壁，风格独特，是上海原有景观风格保持最为完整的老公园。</div></div>','http://20.89.130.136:8080/images/上海中山公园.jpg','http://www.quanjingke.com/jq/m/668'),(15,'世纪公园','<div>位于浦东新区，地铁二、十八号线途经，邻近上海科技馆。大型城市生态公园，国家4A级旅游景区。公园以大 面积草坪、森林、湖泊为主体，体现了东西方园林艺术与自然的融洽，划分为7个区域：湖滨区、观景区、疏林草坪区、鸟类保护区、乡 土田园区、异国园区以及迷你高尔夫球场区。</div></div>','http://20.89.130.136:8080/images/世纪公园.jpg','http://www.quanjingke.com/jq/m/447'),(16,'田子坊','<div>田子坊位于中国上海市泰康路210弄。泰康路是打浦桥地区的一条小街，1998年前这里还是一个马路集市，自1998年9月区政府实施马路集市入室后，把泰康路的路面进行重新铺设，使原来下雨一地泥、天晴一片尘的马路焕然一新。曾经拥挤平常的田 子坊弄堂抹上了苏荷SOHO的色彩，宛如上海市著名的荣乐东路一样，变身成为现代创意聚集地，增添了人文艺术气息。国家3A级别景区，由上海最具特色的石库门里弄演变而来，是时尚地标性创业产业聚集区</div></div>','http://20.89.130.136:8080/images/田子坊.jpg','http://www.quanjingke.com/jq/m/1117'),(17,'朱家角古镇','<div>位于青浦区，地铁十七号线途经。东靠虹桥国际机场，北连昆山，南接嘉兴，西通平望，淀山湖下游、黄金 水道漕港河穿镇而过。镇内河港纵横，九条长街沿河而伸，千栋明清建筑依水而立，36座石桥古风犹存，名胜古迹比比皆是。1991年被上海市政府命名为首批四大文化名镇之一。</div></div>','http://20.89.130.136:8080/images/朱家角古镇.jpg','http://www.quanjingke.com/jq/m/717'),(18,'南京路步行街','<div>位于黄浦区，素有“中华商业第一街”之誉，东起外滩，西至静安寺与延安西路交汇，全长5.5公里，两侧商厦鳞次栉比，云集着约600多家商店，各地的名、特、优、新产品，以及进口的名牌商品，不下数十万种。几家老字号特色商店的商品，尤为名声卓著。</div></div>','http://20.89.130.136:8080/images/南京路步行街.jpg','http://www.quanjingke.com/jq/m/340'),(19,'南园公园','<div>系闽南同乡会泉漳会馆旧址，曾是中国共产党的地下活动场所，内有各种建筑及茔地。</div></div>','http://20.89.130.136:8080/images/南园公园.jpg','http://www.quanjingke.com/jq/m/350'),(20,'上海博物馆','<div>位于上海市黄浦区人民大道201号，陈列面积共计12000平方米，一楼为中国古代青铜馆、中国古代雕塑馆和展览大厅；二楼为中国古代陶瓷馆、暂得楼陶瓷馆和展览厅；三楼为中国历代书法馆、中国历代绘画馆、中国历代玺印馆；四楼为中国古代玉器馆、中国历代钱币馆、中国明清家具馆、中国少数民族工艺馆和展览厅，是综合性博物馆。</div></div>','http://20.89.130.136:8080/images/上海博物馆.jpg','http://www.quanjingke.com/jq/m/407'),(21,'上海城市规划馆','<div>位于上海市黄浦区人民大道100号，靠近人民公园、上海博物馆。地下1层为临展厅，地上1层是序厅，地上1层和2层之间的M层为文创区、阅览区和咖啡馆，地上2、3、4层分别为“人文之城”“创新之城”和“生态之城”展区，5楼设多媒体展示厅。	</div></div>','http://20.89.130.136:8080/images/上海城市规划馆.jpg','http://www.quanjingke.com/jq/m/408'),(22,'滨江大道','<div>位于上海市浦东新区，1997年建成，全长2500米，从泰东路沿黄浦江一直到东昌路，与浦西外滩隔江相望，是集观光、绿化、交通及服务设施为一体的沿江景观工程。它由亲水平台、坡地绿化、半地下厢体及景观道路等组成，被人们赞誉为浦东的新外滩。</div></div>','http://20.89.130.136:8080/images/滨江大道.jpg','http://www.quanjingke.com/jq/m/55'),(23,'上海迪士尼乐园','<div>上海迪士尼乐园，位于上海市浦东新区川沙镇黄赵路310号，于2016年6月16日正式开园，是中国内地首座迪士尼主题乐园，也是中国规模最大的现代服务业中外合作项目之一，是一座具有纯正迪士尼风格并融汇了中国风的主题乐园。上海迪士尼乐园占地1.16平方公里，主题园区分为米奇大街、奇想花园、探险岛、宝藏湾、明日世界、梦幻世界、迪士尼·皮克斯玩具总动员。拥有迪士尼城堡、漫威英雄总部、巴斯光年星际营救等游乐项目。2020年游客量达到550万人次。 2017年，包含上海迪士尼乐园在内的上海迪士尼度假区获“2016—2017年度中国建设工程鲁班奖”。 2020年3月24日，第十一届中国最佳文化旅游大奖榜单发布，上海迪士尼乐园荣获“年度最佳文化旅游大奖”。</div></div><div>很久以前，华特迪士尼公司创始人华特迪士尼先生梦想建造出比普通游乐园更令人流连忘返的场所。“那时两个女儿还小，一般星期六是父亲日，我会带她们去坐旋转木马，到处玩耍。我坐在旁边看时，突然有个想法，我觉得应该建造一个能让父母和子女共度欢乐时光的场所。最终这个想法促成了迪士尼乐园的诞生。但一切始于一位爸爸冥思苦想应该带两个女儿到哪里玩耍，才能一起享受快乐时光。因此，迪士尼乐园诞生于60多年前，致力于让人们回归纯真，感受快乐，分享难忘经历。如今，上海迪士尼度假区以新颖特别的方式延续着这种精神，打造了“探险岛”、“宝藏湾” 和“奇想花园”等全新的主题园区，以及众多独一无二的游乐项目和体验。包括全球迪士尼主题乐园中最大的城堡。华特迪士尼公司的梦想在上海迪士尼度假区延续，欣欣向荣。与任何其他迪士尼产业有所不同的是，上海迪士尼度假区提供独特的、具有中国特色的体验，为形形色色的游客而设计。一切都始于华特先生的想法：建造“一个供父母与子女一起玩乐的家庭乐园。”</div>','http://20.89.130.136:8080/images/上海迪士尼乐园.jpg','暂无网址');
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
INSERT INTO `sites_tags` VALUES (1,1,0.5),(1,2,0.2),(2,3,0.7),(3,1,0.8),(2,2,1),(3,3,1.2),(4,1,NULL),(5,2,NULL),(6,3,NULL),(7,4,NULL),(8,5,NULL),(9,1,NULL),(10,2,NULL),(11,3,NULL),(12,4,NULL),(13,5,NULL),(14,1,NULL),(15,2,NULL),(16,3,NULL),(17,4,NULL),(18,5,NULL),(19,1,NULL),(20,2,NULL),(21,3,NULL),(22,4,NULL),(23,5,NULL);
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
INSERT INTO `tags` VALUES (1,'红色团日'),(2,'自然观光'),(3,'人文博览'),(4,'拓展运动'),(5,'青春游玩');
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
INSERT INTO `users` VALUES ('admin','123456','123123123','http://localhost:8080/images/1676190603.jpg','2023-02-06 20:39:39'),('NP123','Abc12345678','38424@sina.com','http://20.89.130.136:8080/images/3.jpg','2023-02-06 20:54:09'),('测试用户','123456','38424@sina.com','http://20.89.130.136:8080/images/1675838784.jpg','2023-02-08 14:46:24'),('test','123456','38424@sina.com','http://20.89.130.136:8080/images/1675928760.jpg','2023-02-09 15:46:00'),('admin2','123456','38424@sina.com','http://20.89.130.136:8080/images/1675929224.jpg','2023-02-09 15:53:45'),('npneo','1','1','http://20.89.130.136:8080/images/1675929284.jpg','2023-02-09 15:54:45'),('zzz','123456','np123greatest@gmail.com','http://localhost:8080/images/1676188765.jpg','2023-02-12 15:59:26');
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

-- Dump completed on 2023-02-12 16:41:33
