from flask import Flask
import time

app = Flask(__name__)

beginning = None


@app.route('/start')
def benchmark_start():
    global beginning
    beginning = time.time()
    return f"Benchmark start at {beginning}"


@app.route('/stop/<app_name>')
def benchmark_stop(app_name):
    global beginning
    stop = time.time()
    assert beginning is not None
    difference = stop - beginning
    beginning = None
    with open(f"{app_name}.csv", "a") as file:
        file.write(str(difference) + "\n")

    return f"Benchmark stop at {stop}\nDifference: {difference}"
