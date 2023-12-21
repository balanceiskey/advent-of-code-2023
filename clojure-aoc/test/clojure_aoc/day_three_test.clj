(ns clojure-aoc.day-three-test
  (:require [clojure.test :refer :all]
            [clojure-aoc.day-three :refer :all]))


(deftest test-process-match-tuple
  (is (= {:row 0 :length 3 :digits-as-str "467" :digits 467 :idx 0}
         (process-match-tuple 0 ["467" 0]))))

(deftest test-max-out
  (is (= 9 (max-out 25 9))))