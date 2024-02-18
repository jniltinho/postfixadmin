-- DROP TABLE IF EXISTS admin;
CREATE TABLE IF NOT EXISTS admin (
  username varchar(255) COLLATE latin1_general_ci NOT NULL,
  `password` varchar(255) COLLATE latin1_general_ci NOT NULL,
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  superadmin tinyint(1) NOT NULL DEFAULT '0',
  phone varchar(30) CHARACTER SET utf8 NOT NULL DEFAULT '',
  email_other varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
  token varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
  token_validity datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  PRIMARY KEY (username)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Virtual Admins';


-- DROP TABLE IF EXISTS alias;
CREATE TABLE IF NOT EXISTS alias (
  address varchar(255) COLLATE latin1_general_ci NOT NULL,
  goto text COLLATE latin1_general_ci NOT NULL,
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (address),
  KEY domain (domain)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Virtual Aliases';


-- DROP TABLE IF EXISTS alias_domain;
CREATE TABLE IF NOT EXISTS alias_domain (
  alias_domain varchar(255) COLLATE latin1_general_ci NOT NULL DEFAULT '',
  target_domain varchar(255) COLLATE latin1_general_ci NOT NULL DEFAULT '',
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (alias_domain),
  KEY active (active),
  KEY target_domain (target_domain)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Domain Aliases';


-- DROP TABLE IF EXISTS config;
CREATE TABLE IF NOT EXISTS config (
  id int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '',
  `value` varchar(20) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='PostfixAdmin settings';


-- DROP TABLE IF EXISTS domain;
CREATE TABLE IF NOT EXISTS domain (
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  description varchar(255) CHARACTER SET utf8 NOT NULL,
  aliases int(10) NOT NULL DEFAULT '0',
  mailboxes int(10) NOT NULL DEFAULT '0',
  maxquota bigint(20) NOT NULL DEFAULT '0',
  quota bigint(20) NOT NULL DEFAULT '0',
  transport varchar(255) COLLATE latin1_general_ci NOT NULL,
  backupmx tinyint(1) NOT NULL DEFAULT '0',
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  password_expiry int(11) DEFAULT '0',
  PRIMARY KEY (domain)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Virtual Domains';


-- DROP TABLE IF EXISTS domain_admins;
CREATE TABLE IF NOT EXISTS domain_admins (
  username varchar(255) COLLATE latin1_general_ci NOT NULL,
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  id int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (id),
  KEY username (username)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Domain Admins';


-- DROP TABLE IF EXISTS fetchmail;
CREATE TABLE IF NOT EXISTS fetchmail (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  mailbox varchar(255) COLLATE latin1_general_ci NOT NULL,
  src_server varchar(255) COLLATE latin1_general_ci NOT NULL,
  src_auth enum('password','kerberos_v5','kerberos','kerberos_v4','gssapi','cram-md5','otp','ntlm','msn','ssh','any') CHARACTER SET latin1 DEFAULT NULL,
  src_user varchar(255) COLLATE latin1_general_ci NOT NULL,
  src_password varchar(255) COLLATE latin1_general_ci NOT NULL,
  src_folder varchar(255) COLLATE latin1_general_ci NOT NULL,
  poll_time int(11) unsigned NOT NULL DEFAULT '10',
  fetchall tinyint(1) unsigned NOT NULL DEFAULT '0',
  keep tinyint(1) unsigned NOT NULL DEFAULT '0',
  protocol enum('POP3','IMAP','POP2','ETRN','AUTO') CHARACTER SET latin1 DEFAULT NULL,
  usessl tinyint(1) unsigned NOT NULL DEFAULT '0',
  extra_options text COLLATE latin1_general_ci,
  returned_text text COLLATE latin1_general_ci,
  mda varchar(255) COLLATE latin1_general_ci NOT NULL,
  `date` timestamp NOT NULL DEFAULT '2000-01-01 00:00:00',
  sslcertck tinyint(1) NOT NULL DEFAULT '0',
  sslcertpath varchar(255) CHARACTER SET utf8 DEFAULT '',
  sslfingerprint varchar(255) COLLATE latin1_general_ci DEFAULT '',
  domain varchar(255) COLLATE latin1_general_ci DEFAULT '',
  active tinyint(1) NOT NULL DEFAULT '0',
  created timestamp NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  src_port int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;


-- DROP TABLE IF EXISTS log;
CREATE TABLE IF NOT EXISTS log (
  `timestamp` datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  username varchar(255) COLLATE latin1_general_ci NOT NULL,
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  `action` varchar(255) COLLATE latin1_general_ci NOT NULL,
  `data` text COLLATE latin1_general_ci NOT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (id),
  KEY `timestamp` (`timestamp`),
  KEY domain_timestamp (domain,`timestamp`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Log';


-- DROP TABLE IF EXISTS mailbox;
CREATE TABLE IF NOT EXISTS mailbox (
  username varchar(255) COLLATE latin1_general_ci NOT NULL,
  `password` varchar(255) COLLATE latin1_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  maildir varchar(255) COLLATE latin1_general_ci NOT NULL,
  quota bigint(20) NOT NULL DEFAULT '0',
  local_part varchar(255) COLLATE latin1_general_ci NOT NULL,
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  modified datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  phone varchar(30) CHARACTER SET utf8 NOT NULL DEFAULT '',
  email_other varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
  token varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
  token_validity datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  password_expiry datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  PRIMARY KEY (username),
  KEY domain (domain)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Virtual Mailboxes';


-- DROP TABLE IF EXISTS quota;
CREATE TABLE IF NOT EXISTS quota (
  username varchar(255) COLLATE latin1_general_ci NOT NULL,
  path varchar(100) COLLATE latin1_general_ci NOT NULL,
  `current` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (username,path)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci;


-- DROP TABLE IF EXISTS quota2;
CREATE TABLE IF NOT EXISTS quota2 (
  username varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL,
  bytes bigint(20) NOT NULL DEFAULT '0',
  messages int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (username)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- DROP TABLE IF EXISTS vacation;
CREATE TABLE IF NOT EXISTS vacation (
  email varchar(255) COLLATE latin1_general_ci NOT NULL,
  `subject` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  body text CHARACTER SET utf8mb4 NOT NULL,
  `cache` text COLLATE latin1_general_ci NOT NULL,
  domain varchar(255) COLLATE latin1_general_ci NOT NULL,
  created datetime NOT NULL DEFAULT '2000-01-01 00:00:00',
  active tinyint(1) NOT NULL DEFAULT '1',
  modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  activefrom timestamp NOT NULL DEFAULT '2000-01-01 00:00:00',
  activeuntil timestamp NOT NULL DEFAULT '2038-01-18 00:00:00',
  interval_time int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (email),
  KEY email (email)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_general_ci COMMENT='Postfix Admin - Virtual Vacation';


-- DROP TABLE IF EXISTS vacation_notification;
CREATE TABLE IF NOT EXISTS vacation_notification (
  on_vacation varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL,
  notified varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '',
  notified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (on_vacation,notified),
  CONSTRAINT vacation_notification_pkey FOREIGN KEY (on_vacation) REFERENCES vacation (email) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Postfix Admin - Virtual Vacation Notifications';

