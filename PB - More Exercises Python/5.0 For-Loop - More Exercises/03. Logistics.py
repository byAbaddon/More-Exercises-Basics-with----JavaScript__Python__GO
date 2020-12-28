cargo_count = int(input())
cargos_dict = {'all_cargo' : 0, 'bus': 0, 'truck' : 0, 'train' : 0}

for _ in range(cargo_count):
    current_cargo = int(input())
    if current_cargo <= 3:
        cargos_dict['bus'] += current_cargo
    elif 4 <= current_cargo <= 11:
        cargos_dict['truck'] += current_cargo
    else:
        cargos_dict['train'] += current_cargo
    cargos_dict['all_cargo'] += current_cargo

average_price = (cargos_dict['bus'] * 200 + cargos_dict['truck'] * 175 + cargos_dict['train'] * 120) / cargos_dict['all_cargo']

print(f"{average_price:.2f}")
print(f"{cargos_dict['bus'] / cargos_dict['all_cargo'] *  100:.2f}%")
print(f"{cargos_dict['truck'] / cargos_dict['all_cargo'] * 100:.2f}%")
print(f"{cargos_dict['train'] / cargos_dict['all_cargo'] * 100:.2f}%")


'''
4
1
5
16
3
'''
