period = int(input())

counter = 0
untreated_patients = 0
treated_patients = 0
doctors = 7

for patient in range(period):
    number_of_patient = int(input())
    counter += 1
    if counter % 3 == 0:
        if untreated_patients > treated_patients:
            doctors += 1
    if 0 <= number_of_patient <= doctors:
        treated_patients += number_of_patient
    else:
        untreated_patients = (number_of_patient - doctors) + untreated_patients
        treated_patients += doctors


print(f"Treated patients: {treated_patients}.\nUntreated patients: {untreated_patients}.")
