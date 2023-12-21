(ns clojure-aoc.day-three
  (:require [clojure-aoc.utils :as utils]))

(def puzzle-ex (utils/get-rows "day-three-ex.txt"))
(def puzzle-pt-1 (utils/get-rows "day-three-part-one.txt"))

(defn zero-out
  "If number is less than zero, returns zero."
  [num] (if (< num 0) 0 num))

(defn max-out
  "If number exceeds max, returns max."
  [val max] (if (> val max) max val))

; dear god, I wish there was a better way to do this.
(defn find-numbers
  "Provided an input string like '...342...23' returns [['342' 3] ['23' 11]]"
  [input-str]
  (let [matcher (re-matcher #"\d+" input-str)]
    (loop [matches []]
      (if-let [match (re-find matcher)]
        (let [position (.start matcher)
              match-with-pos [match position]]
          (recur (conj matches match-with-pos)))
        matches))))

(defn find-gears
  [input-str]
  (let [matcher (re-matcher #"\*" input-str)]
    (loop [matches []]
      (if-let [match (re-find matcher)]
        (let [position (.start matcher)
              match-with-pos [match position]]
          (recur (conj matches match-with-pos)))
        matches))))
(defn get-potential-numbers [rows]
  (map #(find-numbers %) rows))

(defn get-potential-gears [rows]
  (map #(find-gears %) rows))

(defn contains-symbol? [input] (boolean (re-find #"[^0-9.]" input)))
(defn process-match-tuple [row-index [digits idx]]
  {:row    row-index
   :length (count digits)
   :digits-as-str digits
   :digits (Integer/parseInt digits)
   :idx    idx})

(defn process-matches
  "Takes a list of matches ([[\"467\" 0] [\"114\" 5]] [[\"35\" 2]])
  and returns a list of structured maps of the form {:row 0 :length 2 :digits \"46\" :idx 2}"
  [matches]
  (mapcat (fn [row-index row]
            (map (partial process-match-tuple row-index) row))
          (range)
          matches))

(defn get-search-substr [puzzle {row-idx :row str-start :idx :keys [length]}]
  (let [start (zero-out (dec str-start))
        end (max-out (+ 1 str-start length) (dec (count (get puzzle row-idx))))
        input-row-top (get puzzle (zero-out (dec row-idx)))
        input-row-middle (get puzzle row-idx)
        input-row-bottom (get puzzle (max-out (inc row-idx) (dec (count puzzle))))]
    [
     (subs input-row-top start end)
     (subs input-row-middle start end)
     (subs input-row-bottom start end)]))


(defn get-gear-search-box [puzzle {row-idx :row str-start :idx :keys [length]}]
  (let [start (zero-out (dec str-start))
        end (max-out (+ 1 str-start length) (dec (count (get puzzle row-idx))))
        input-row-top (get puzzle (zero-out (dec row-idx)))
        input-row-middle (get puzzle row-idx)
        input-row-bottom (get puzzle (max-out (inc row-idx) (dec (count puzzle))))]

    [{:row-idx (zero-out (dec row-idx)) :substr (subs input-row-top start end) :start start}
     {:row-idx row-idx :substr (subs input-row-middle start end) :start start}
     {:row-idx (max-out (inc row-idx) (dec (count puzzle))) :substr (subs input-row-bottom start end) :start start}]))

(defn number-near-symbol? [puzzle row-obj]
  (some contains-symbol? (get-search-substr puzzle row-obj)))

(defn enhance [{:keys [row-idx substr start]}]
  (let [gear-idx (.indexOf substr "*")]
    (if (> gear-idx -1)
      [row-idx (+ start gear-idx)]
      nil)))

;; I gotta think about this more, my plan is falling apart because
; substrings will tell you what to search in, _but_ it breaks down because
; you then need to have consistent coordinates for where a gear might be located and that's relative
; to the entire row/puzzle. Getting the row index for a gear is right, getting it's position
; in the row is wrong. I think there's some math here. We can maybe
(defn enhance-with-gear [puzzle row-obj]
  (assoc row-obj :gears (vec (filter some? (map enhance (get-gear-search-box puzzle row-obj))))))

(number-near-symbol?
  puzzle-ex
  {:row 9, :length 3, :digits "598", :idx 5})

(process-matches (get-potential-numbers puzzle-ex))

(defn eval-puzzle-input [input]
  (reduce + (map :digits (filter
    (partial number-near-symbol? input)
    (process-matches (get-potential-numbers input))))))

(process-matches (get-potential-numbers puzzle-ex))

(defn eval-puzzle-input-pt-2 [input]
  (map
    (partial enhance-with-gear input)
    (process-matches (get-potential-numbers input))))


(defn get-gears [structs]
  (set (filter some? (mapcat (fn [x] (if (empty? (:gears x))
                 nil
                 (:gears x))
         ) structs))))

(get-potential-gears puzzle-ex)
(eval-puzzle-input puzzle-ex)
(eval-puzzle-input puzzle-pt-1)
(eval-puzzle-input-pt-2 puzzle-ex)
(eval-puzzle-input-pt-2 puzzle-ex)

(defn has-gear? [gear-coords {:keys [gears]}]
  (some #(= gear-coords %) gears))
;(has-gear? [1 3] (first (eval-puzzle-input-pt-2 puzzle-ex)))

(defn get-multiple [structs gear-coords]
  (let [matches (filter (partial has-gear? gear-coords) structs)
        digits (map :digits matches)]
    (if (> (count matches) 1) (apply * digits) 0)))


; example (swap puzzle-pt-1, puzzle-ex etc here)
(let [results (eval-puzzle-input-pt-2 puzzle-pt-1)
      uniq-gears (get-gears results)]
  (reduce + (map (partial get-multiple results) uniq-gears)))

;















