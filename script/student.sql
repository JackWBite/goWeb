CREATE TABLE `student` (
  `sid` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL COMMENT '学生名称',
  `age` varchar(45) DEFAULT NULL COMMENT '年龄',
  `class` varchar(45) DEFAULT NULL COMMENT '班级',
  `school` varchar(45) DEFAULT NULL COMMENT '学校',
  `sex` int DEFAULT NULL COMMENT '性别，1:男 2:女',
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='学生表'