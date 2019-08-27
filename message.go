package main

func queroCafeMessage() string {
	msg := `Semana passada compartilhei um post sobre movimentos sociais, um claro shitpost de uma página grande, teve poucas curtidas, mas até aí tudo ok
O problema começa quando na segunda feira recebo um email do RH do meu trabalho dizendo que precisava conversar comigo com urgência, marca a reunião pra quinta (ontem) e eu sem saber o que tava acontecendo
Chego na reunião e ele tem vários prints de comentários meus aqui e da publicação
Me conta que alguém tinha enviado um email e pedindo a abertura processo contra mim por discurso de ódio na internet e um posicionamento da coordenação do curso sobre o caso
Meu orientador disse que tentou abafar o caso e conversar comigo primeiro e entender o que tinha acontecido, mas me disse que tinham enviado a mesma coisa pro chefe e que ele teria que tomar uma providência sobre o caso
Foi ligado para minha família e ao que tudo indica eu posso perder minha vaga no emprego
Além disso tudo saíram os prints no spotted da empresa com o nome borrado, e, claro, fora de contexto parecem coisas horríveis
A empresa toda ta com ódio mortal do dono dos comentários e eu to temendo por minha segurança
Minha irmã é amiga de um advogado que ta ajudando a gente, ele entrou em contato com a pessoa que enviou os emails e descobriu que ele queria comer o cu de quem tá lendo.`

	return msg
}

func helpMessage() string {
	msg := `Comandos disponíveis:
/start - Você será notificado automaticamente assim que o restaurante abrir
/cancel - Cancela a notificação automática
/status - Verifica se o restaurante está aberto ou não
/help - Mostra esta mensagem
Para mais informações veja a descrição do bot.
@marmitaz_bot`

	return msg
}
