anterior = 0
proximo = 0
posicao = 0

while(proximo < 100000000000000):
  posicao += 1
  print(proximo)
  print(f'Posição: {posicao}')
  proximo = proximo + anterior
  anterior = proximo - anterior
  if(proximo == 0):
    proximo = proximo + 1
