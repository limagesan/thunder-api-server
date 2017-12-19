#!/bin/sh

if [ "$1" == "clean" ]
then
  echo "called: clean"
  array=()
  for file in `\find ./database -maxdepth 1 -type f`; do
    # echo $file
    array=("${array[@]}" $file) 
  done

  max=0
  i=0
  for e in ${array[@]}; do
    # echo "array[$i]=${e}"
    
    num=`echo $e | sed -e 's/[^0-9]//g'`
    # echo "num is: ${num}"
    if [ "$num" != "" ]
    then  
      if [ "$num" -gt "$max" ]
      then
        max=$num
      fi
      let i++
    fi
  done
  # echo "max num is ${max}"
  indexNum=`expr $max + 1`
  fileName="thunder${indexNum}.db"
  if test -e "./database/thunder.db"
  then
    mv ./database/thunder.db ./database/$fileName
    echo "SUCCESS: move thunder.db to ${fileName}"
  else
    echo "EXCEPTION: ./database/thunder.db is not found"
  fi

elif [ "$1" == "cp" ]
then
  cp ../scraper-thunder/database/thunder.db ./database/
  echo "SUCCESS: copy from scraper-thunder"
  ls database
else 
  echo "unknown command is called"
fi
