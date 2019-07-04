init:
	python3  -m  venv  env; \
	echo 'source env/bin/activate' >> .env; \
	source ./env/bin/activate; \
	pip3 install -r requirements.txt; \
	pip3 install -U pytest; 

run:
	python integration_test.py

tests:
	pytest -o log_cli=true *.py

update_deps:
	source env/bin/activate; \
	pip3 install --upgrade -r requirements.txt
