-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- 数据库： `service_items_go`
--

-- --------------------------------------------------------

--
-- 表的结构 `items`
--

CREATE TABLE `items` (
    `id` int(11) NOT NULL,
    `appkey` varchar(64) NOT NULL,
    `channel` int(11) NOT NULL,
    `item_id` varchar(64) NOT NULL COMMENT '商品ID',
    `name` varchar(255) NOT NULL COMMENT '商品名称',
    `photo` varchar(512) NOT NULL COMMENT '商品主图',
    `detail` text NOT NULL COMMENT '商品详情',
    `state` tinyint(4) NOT NULL COMMENT '商品状态 0 正常 1 已删除 2 已彻底删除',
    `updated_at` datetime NOT NULL,
    `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- --------------------------------------------------------

--
-- 表的结构 `item_searches`
--

CREATE TABLE `item_searches` (
    `id` int(11) NOT NULL,
    `appkey` varchar(64) NOT NULL,
    `channel` int(11) NOT NULL,
    `item_id` varchar(64) NOT NULL COMMENT '商品ID',
    `sku_id` varchar(64) NOT NULL COMMENT '商品SKU_ID',
    `item_name` varchar(255) NOT NULL COMMENT '商品名称',
    `sku_name` varchar(255) NOT NULL COMMENT '商品SKU名称',
    `barcode` varchar(50) NOT NULL COMMENT '条形码',
    `item_state` tinyint(4) NOT NULL COMMENT '商品状态 0 正常 1 已删除 2 已彻底删除',
    `sku_state` tinyint(4) NOT NULL COMMENT 'sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除',
    `updated_at` datetime NOT NULL,
    `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品搜索表';

-- --------------------------------------------------------

--
-- 表的结构 `skus`
--

CREATE TABLE `skus` (
    `id` int(11) NOT NULL,
    `appkey` varchar(64) NOT NULL,
    `channel` int(11) NOT NULL,
    `item_id` varchar(64) NOT NULL COMMENT '商品ID',
    `sku_id` varchar(64) NOT NULL COMMENT '商品SKU_ID',
    `name` varchar(255) NOT NULL COMMENT '商品SKU名称',
    `photo` varchar(512) NOT NULL COMMENT '商品SKU主图',
    `barcode` varchar(50) NOT NULL COMMENT '条形码',
    `state` tinyint(4) NOT NULL COMMENT 'sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除',
    `updated_at` datetime NOT NULL,
    `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `items`
--
ALTER TABLE `items`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `item` (`appkey`,`channel`,`item_id`);

--
-- 表的索引 `item_searches`
--
ALTER TABLE `item_searches`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `item_search` (`appkey`,`channel`,`item_id`,`sku_id`);

--
-- 表的索引 `skus`
--
ALTER TABLE `skus`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `sku` (`appkey`,`channel`,`item_id`,`sku_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `items`
--
ALTER TABLE `items`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `item_searches`
--
ALTER TABLE `item_searches`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `skus`
--
ALTER TABLE `skus`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back