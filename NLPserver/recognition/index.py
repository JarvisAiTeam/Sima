from __future__ import print_function
from nltk.tag import DefaultTagger
from nltk.tag import UnigramTagger
from nltk.tag import AffixTagger
from  nltk import word_tokenize
from nltk.tag import tnt
from nltk.tag.sequential import ClassifierBasedPOSTagger
from nltk.classify import MaxentClassifier
import codecs
import numpy as np
import pickle
import os


corpus_path = 'C:/Users/niksy/PycharmProjects/NLPserver/corpus/pos_tagger_corpus.dat'
max_sent = 1000000


def train_speech_recognition():
    global unigram_tagger

    print('Loading "{0}...'.format(corpus_path))
    corpus = []
    ntoken = 0
    n_bad = 0
    with codecs.open(corpus_path, 'r', 'utf-8') as rdr:
        sent = []
        good = True
        for line0 in rdr:
            line = line0.strip()
            if len(line) == 0:
                if good:
                    corpus.append(sent)
                    ntoken += len(sent)
                    if len(corpus) >= max_sent:
                        break
                else:
                    n_bad += 1
                good = True
                sent = []
            else:
                tx = line.split('\t')
                if len(tx) < 2:
                    good = False
                else:
                    word = tx[0].lower()
                    pos = tx[1].lower()
                    sent.append((word, pos))

    print('done, {0} good sentences, {1} ntoken'.format(len(corpus), ntoken))
    # ----------------------------------------------------------------------

    n_patterns = len(corpus)

    n_test = int(n_patterns * 0.1)
    n_train = n_patterns - n_test
    print('n_test={0} n_train={1}'.format(n_test, n_train))
    data_indeces = [x for x in range(n_patterns)]
    np.random.shuffle(data_indeces)
    test_indeces = data_indeces[: n_test]
    train_indeces = data_indeces[n_test:]

    train_corpus = [corpus[i] for i in train_indeces]
    test_corpus = [corpus[i] for i in test_indeces]

    # ----------------------------------------------------------------------

    default_tagger = DefaultTagger(u'РЎРЈР©Р•РЎРўР’РРўР•Р›Р¬РќРћР•')

    # ----------------------------------------------------------------------

    # print( 'Training AffixTagger on 1-suffixes...' )
    suffix1_tagger = AffixTagger(train_corpus, affix_length=-1, backoff=default_tagger)
    # print( 'Testing...' )
    acc = suffix1_tagger.evaluate(test_corpus)
    # print( 'AffixTagger(1) accuracy={0}\n'.format(acc) )

    # ----------------------------------------------------------------------

    # print( 'Training AffixTagger on 2-suffixes...' )
    suffix2_tagger = AffixTagger(train_corpus, affix_length=-2, backoff=suffix1_tagger)
    # print( 'Testing...' )
    acc = suffix2_tagger.evaluate(test_corpus)
    # print( 'AffixTagger(2,1) accuracy={0}\n'.format(acc) )

    # ----------------------------------------------------------------------

    # print( 'Training AffixTagger on 3-suffixes...' )
    suffix3_tagger = AffixTagger(train_corpus, affix_length=-3, backoff=suffix2_tagger)
    # print( 'Testing...' )
    acc = suffix3_tagger.evaluate(test_corpus)
    # print( 'AffixTagger(3,2,1) accuracy={0}\n'.format(acc) )

    # ----------------------------------------------------------------------

    # print( 'Training AffixTagger on 4,3,2-suffixes...' )
    suffix4_tagger = AffixTagger(train_corpus, affix_length=-4, backoff=suffix3_tagger)
    # print( 'Testing...' )
    acc = suffix4_tagger.evaluate(test_corpus)
    # print( 'AffixTagger(4,3,2) accuracy={0}\n'.format(acc) )

    # ----------------------------------------------------------------------

    # print( 'Testing UnigramTagger + AffixTagger(4,3,2,1)...' )
    unigram_tagger = UnigramTagger(train_corpus, backoff=suffix4_tagger)

    # print(unigram_tagger.tag(word_tokenize("погода на завтра в Одессе")))
    acc = unigram_tagger.evaluate(test_corpus)
    # print( 'UnigramTagger+AffixTagger(4,3,2,1) accuracy={0}\n'.format(acc) )
    cache_model()

# ----------------------------------------------------------------------


def cache_model():
    with open("trained_model.file", "wb") as f:
        pickle.dump(unigram_tagger, f, pickle.HIGHEST_PROTOCOL)
# ----------------------------------------------------------------------


def check_cache():
    global unigram_tagger

    if os.path.exists("trained_model.file"):
        with open("trained_model.file", "rb") as f:
            dump = pickle.load(f)
        if dump is not None:
            unigram_tagger = dump
            return True
        else:
            return False
    return False


def tag_sentence(sentence):
    return unigram_tagger.tag(word_tokenize(sentence.lower()))



