<!DOCTYPE html>
<html lang="en"><head></head>
<body>
	<form action="/image" method="post">
		<label>Digite seu nome</label>
		<input type="text" name="nome" id="nome"><br/><br/>

		<label>Taxa de crossover</label>
		<input type="text" name="taxaCrossover" id="taxaCrossover" value="0.8"><br/><br/>

		<label>Taxa de mutação</label>
		<input type="text" name="taxaMutacao" id="taxaMutacao" value="0.3"><br/><br/>

		<label>Elitismo</label><br/>
		Sim <input type="radio" name="elitismo" value="S" checked="checked"><br/>
		Não <input type="radio" name="elitismo" value="N"><br/><br/>

		<label>Tamanho da população</label>
		<input type="text" name="tamanhoPopulacao" id="tamanhoPopulacao" value="2000"><br/>

		<br/>
		<input type="submit" value="Enviar">
	</form>
</body>
</html>