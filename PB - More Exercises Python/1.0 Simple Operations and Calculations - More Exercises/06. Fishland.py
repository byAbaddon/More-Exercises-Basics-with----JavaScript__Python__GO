skumriq_price, caca_price, kgPalamud, kgSafrid, kgMidi  = [float(input()) for _ in range(5)]
palamud_price = skumriq_price + 0.6 * skumriq_price
safrid_price = caca_price + 0.8 * caca_price
price = (kgPalamud * palamud_price + kgSafrid * safrid_price + kgMidi * 7.5)
print(f'{price:.2f}')


'''
6.90
4.20
1.5
2.5
1
'''
