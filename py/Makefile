init:
	python3 -m venv env; \
	echo 'source venv/bin/activate' >> .env; \
	source ./env/bin/activate; \
	pip3 install -r requirements.txt; \
	pip install -U pytest \

run:
	python main.py

test:
	pytest test.py

update_deps:
	source env/bin/activate; \
	pip3 install --upgrade -r requirements.txt
