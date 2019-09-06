go build editor.go

echo "----------------------------First Round----------------------------"
(time ./editor csvs/csv1/csv_file_1.csv) 2>> seqTime1.txt
echo "1seq"

(time ./editor csvs/csv2/csv_file_2.csv) 2>> seqTime2.txt
echo "2seq"

(time ./editor csvs/csv3/csv_file_3.csv) 2>> seqTime3.txt
echo "3seq"

(time ./editor -p=1 csvs/csv1/csv_file_1.csv) 2>> para1-1.txt
echo "1para1"
(time ./editor -p=2 csvs/csv1/csv_file_1.csv) 2>> para1-2.txt
echo "1para2"
(time ./editor -p=4 csvs/csv1/csv_file_1.csv) 2>> para1-4.txt
echo "1para4"
(time ./editor -p=6 csvs/csv1/csv_file_1.csv) 2>> para1-6.txt
echo "1para6"
(time ./editor -p=8 csvs/csv1/csv_file_1.csv) 2>> para1-8.txt
echo "1para8"

(time ./editor -p=1 csvs/csv2/csv_file_2.csv) 2>> para2-1.txt
echo "2para1"
(time ./editor -p=2 csvs/csv2/csv_file_2.csv) 2>> para2-2.txt
echo "2para2"
(time ./editor -p=4 csvs/csv2/csv_file_2.csv) 2>> para2-4.txt
echo "2para4"
(time ./editor -p=6 csvs/csv2/csv_file_2.csv) 2>> para2-6.txt
echo "2para6"
(time ./editor -p=8 csvs/csv2/csv_file_2.csv) 2>> para2-8.txt
echo "2para8"

(time ./editor -p=1 csvs/csv3/csv_file_3.csv) 2>> para3-1.txt
echo "3para1"
(time ./editor -p=2 csvs/csv3/csv_file_3.csv) 2>> para3-2.txt
echo "3para2"
(time ./editor -p=4 csvs/csv3/csv_file_3.csv) 2>> para3-4.txt
echo "3para4"
(time ./editor -p=6 csvs/csv3/csv_file_3.csv) 2>> para3-6.txt
echo "3para6"
(time ./editor -p=8 csvs/csv3/csv_file_3.csv) 2>> para3-8.txt
echo "3para8"

echo "----------------------------Second Round----------------------------"
(time ./editor csvs/csv1/csv_file_1.csv) 2>> seqTime1.txt
echo "1seq"

(time ./editor csvs/csv2/csv_file_2.csv) 2>> seqTime2.txt
echo "2seq"

(time ./editor csvs/csv3/csv_file_3.csv) 2>> seqTime3.txt
echo "3seq"

(time ./editor -p=1 csvs/csv1/csv_file_1.csv) 2>> para1-1.txt
echo "1para1"
(time ./editor -p=2 csvs/csv1/csv_file_1.csv) 2>> para1-2.txt
echo "1para2"
(time ./editor -p=4 csvs/csv1/csv_file_1.csv) 2>> para1-4.txt
echo "1para4"
(time ./editor -p=6 csvs/csv1/csv_file_1.csv) 2>> para1-6.txt
echo "1para6"
(time ./editor -p=8 csvs/csv1/csv_file_1.csv) 2>> para1-8.txt
echo "1para8"

(time ./editor -p=1 csvs/csv2/csv_file_2.csv) 2>> para2-1.txt
echo "2para1"
(time ./editor -p=2 csvs/csv2/csv_file_2.csv) 2>> para2-2.txt
echo "2para2"
(time ./editor -p=4 csvs/csv2/csv_file_2.csv) 2>> para2-4.txt
echo "2para4"
(time ./editor -p=6 csvs/csv2/csv_file_2.csv) 2>> para2-6.txt
echo "2para6"
(time ./editor -p=8 csvs/csv2/csv_file_2.csv) 2>> para2-8.txt
echo "2para8"

(time ./editor -p=1 csvs/csv3/csv_file_3.csv) 2>> para3-1.txt
echo "3para1"
(time ./editor -p=2 csvs/csv3/csv_file_3.csv) 2>> para3-2.txt
echo "3para2"
(time ./editor -p=4 csvs/csv3/csv_file_3.csv) 2>> para3-4.txt
echo "3para4"
(time ./editor -p=6 csvs/csv3/csv_file_3.csv) 2>> para3-6.txt
echo "3para6"
(time ./editor -p=8 csvs/csv3/csv_file_3.csv) 2>> para3-8.txt
echo "3para8"


echo "----------------------------Third Round----------------------------"
(time ./editor csvs/csv1/csv_file_1.csv) 2>> seqTime1.txt
echo "1seq"

(time ./editor csvs/csv2/csv_file_2.csv) 2>> seqTime2.txt
echo "2seq"

(time ./editor csvs/csv3/csv_file_3.csv) 2>> seqTime3.txt
echo "3seq"

(time ./editor -p=1 csvs/csv1/csv_file_1.csv) 2>> para1-1.txt
echo "1para1"
(time ./editor -p=2 csvs/csv1/csv_file_1.csv) 2>> para1-2.txt
echo "1para2"
(time ./editor -p=4 csvs/csv1/csv_file_1.csv) 2>> para1-4.txt
echo "1para4"
(time ./editor -p=6 csvs/csv1/csv_file_1.csv) 2>> para1-6.txt
echo "1para6"
(time ./editor -p=8 csvs/csv1/csv_file_1.csv) 2>> para1-8.txt
echo "1para8"

(time ./editor -p=1 csvs/csv2/csv_file_2.csv) 2>> para2-1.txt
echo "2para1"
(time ./editor -p=2 csvs/csv2/csv_file_2.csv) 2>> para2-2.txt
echo "2para2"
(time ./editor -p=4 csvs/csv2/csv_file_2.csv) 2>> para2-4.txt
echo "2para4"
(time ./editor -p=6 csvs/csv2/csv_file_2.csv) 2>> para2-6.txt
echo "2para6"
(time ./editor -p=8 csvs/csv2/csv_file_2.csv) 2>> para2-8.txt
echo "2para8"

(time ./editor -p=1 csvs/csv3/csv_file_3.csv) 2>> para3-1.txt
echo "3para1"
(time ./editor -p=2 csvs/csv3/csv_file_3.csv) 2>> para3-2.txt
echo "3para2"
(time ./editor -p=4 csvs/csv3/csv_file_3.csv) 2>> para3-4.txt
echo "3para4"
(time ./editor -p=6 csvs/csv3/csv_file_3.csv) 2>> para3-6.txt
echo "3para6"
(time ./editor -p=8 csvs/csv3/csv_file_3.csv) 2>> para3-8.txt
echo "3para8"


echo "----------------------------Forth Round----------------------------"
(time ./editor csvs/csv1/csv_file_1.csv) 2>> seqTime1.txt
echo "1seq"

(time ./editor csvs/csv2/csv_file_2.csv) 2>> seqTime2.txt
echo "2seq"

(time ./editor csvs/csv3/csv_file_3.csv) 2>> seqTime3.txt
echo "3seq"

(time ./editor -p=1 csvs/csv1/csv_file_1.csv) 2>> para1-1.txt
echo "1para1"
(time ./editor -p=2 csvs/csv1/csv_file_1.csv) 2>> para1-2.txt
echo "1para2"
(time ./editor -p=4 csvs/csv1/csv_file_1.csv) 2>> para1-4.txt
echo "1para4"
(time ./editor -p=6 csvs/csv1/csv_file_1.csv) 2>> para1-6.txt
echo "1para6"
(time ./editor -p=8 csvs/csv1/csv_file_1.csv) 2>> para1-8.txt
echo "1para8"

(time ./editor -p=1 csvs/csv2/csv_file_2.csv) 2>> para2-1.txt
echo "2para1"
(time ./editor -p=2 csvs/csv2/csv_file_2.csv) 2>> para2-2.txt
echo "2para2"
(time ./editor -p=4 csvs/csv2/csv_file_2.csv) 2>> para2-4.txt
echo "2para4"
(time ./editor -p=6 csvs/csv2/csv_file_2.csv) 2>> para2-6.txt
echo "2para6"
(time ./editor -p=8 csvs/csv2/csv_file_2.csv) 2>> para2-8.txt
echo "2para8"

(time ./editor -p=1 csvs/csv3/csv_file_3.csv) 2>> para3-1.txt
echo "3para1"
(time ./editor -p=2 csvs/csv3/csv_file_3.csv) 2>> para3-2.txt
echo "3para2"
(time ./editor -p=4 csvs/csv3/csv_file_3.csv) 2>> para3-4.txt
echo "3para4"
(time ./editor -p=6 csvs/csv3/csv_file_3.csv) 2>> para3-6.txt
echo "3para6"
(time ./editor -p=8 csvs/csv3/csv_file_3.csv) 2>> para3-8.txt
echo "3para8"


echo "----------------------------Last Round----------------------------"
(time ./editor csvs/csv1/csv_file_1.csv) 2>> seqTime1.txt
echo "1seq"

(time ./editor csvs/csv2/csv_file_2.csv) 2>> seqTime2.txt
echo "2seq"

(time ./editor csvs/csv3/csv_file_3.csv) 2>> seqTime3.txt
echo "3seq"

(time ./editor -p=1 csvs/csv1/csv_file_1.csv) 2>> para1-1.txt
echo "1para1"
(time ./editor -p=2 csvs/csv1/csv_file_1.csv) 2>> para1-2.txt
echo "1para2"
(time ./editor -p=4 csvs/csv1/csv_file_1.csv) 2>> para1-4.txt
echo "1para4"
(time ./editor -p=6 csvs/csv1/csv_file_1.csv) 2>> para1-6.txt
echo "1para6"
(time ./editor -p=8 csvs/csv1/csv_file_1.csv) 2>> para1-8.txt
echo "1para8"

(time ./editor -p=1 csvs/csv2/csv_file_2.csv) 2>> para2-1.txt
echo "2para1"
(time ./editor -p=2 csvs/csv2/csv_file_2.csv) 2>> para2-2.txt
echo "2para2"
(time ./editor -p=4 csvs/csv2/csv_file_2.csv) 2>> para2-4.txt
echo "2para4"
(time ./editor -p=6 csvs/csv2/csv_file_2.csv) 2>> para2-6.txt
echo "2para6"
(time ./editor -p=8 csvs/csv2/csv_file_2.csv) 2>> para2-8.txt
echo "2para8"

(time ./editor -p=1 csvs/csv3/csv_file_3.csv) 2>> para3-1.txt
echo "3para1"
(time ./editor -p=2 csvs/csv3/csv_file_3.csv) 2>> para3-2.txt
echo "3para2"
(time ./editor -p=4 csvs/csv3/csv_file_3.csv) 2>> para3-4.txt
echo "3para4"
(time ./editor -p=6 csvs/csv3/csv_file_3.csv) 2>> para3-6.txt
echo "3para6"
(time ./editor -p=8 csvs/csv3/csv_file_3.csv) 2>> para3-8.txt
echo "3para8"

echo "ALL FINISHED!!!!!!!!!!"
