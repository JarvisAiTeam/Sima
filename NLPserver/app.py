from flask import Flask
from flask import jsonify
from flask import request
from recognition import index


app = Flask(__name__, root_path='corpus/')

res = dict()


def train():
    if not index.check_cache():
        index.train_speech_recognition()


train()


@app.route('/recognition/speech_part', methods=['POST'])
def app_sent():
    res["result"] = index.tag_sentence(request.form['sentence'])
    return jsonify(res)


if __name__ == '__main__':
    app.run()



