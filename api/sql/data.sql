insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$nxVbrL0L7BC3rsf5XSFsjOQov2GqCVaZ1nUfSZVKlhMsNqNnExnbO"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$nxVbrL0L7BC3rsf5XSFsjOQov2GqCVaZ1nUfSZVKlhMsNqNnExnbO"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$nxVbrL0L7BC3rsf5XSFsjOQov2GqCVaZ1nUfSZVKlhMsNqNnExnbO"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! OK!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! OK!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! OK!", 3);