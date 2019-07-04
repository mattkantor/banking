init:
	python3  -m  venv  env; \
	echo 'source env/bin/activate' >> .env; \
	source ./env/bin/activate; \
	pip3 install -r requirements.txt; \
	pip3 install -U pytest; 

run:
	python main.py

tests:
	pytest test_code.py

update_deps:
	source env/bin/activate; \
	pip3 install --upgrade -r requirements.txt
