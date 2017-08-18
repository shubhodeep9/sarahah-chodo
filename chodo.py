import csv, requests

def hasSpaceAndAlpha(string):
    return any(char.isalpha() for char in string) and any(char.isspace() for char in string) and all(char.isalpha() or char.isspace() for char in string)

data = []

with open('Indian-Male-Names.csv') as file:
	reader = csv.DictReader(file)
	for row in reader:
		if hasSpaceAndAlpha(row['name']):
			row['name'] = row['name'].replace(' ','')
			data.append(row['name'])

selected = []
try:
	for item in data:
		r = requests.get('https://'+item+'.sarahah.com', allow_redirects=False)
		print (item, len(selected))
		if r.status_code == 200:
			selected.append(item)
except Exception as e:
	print(e.args[0])
	pass

print(len(selected))
