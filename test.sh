FILES=$(ls tests/)

for file in $FILES; {
    echo running $file
    ./minigo run tests/$file > /dev/null 2>&1
    execution_status=$?
    type=$(echo $file | rg -q error)
    if [[ $execution_status -eq $? ]]; then
        echo unexpected status while running $file check its contents
        exit
    fi
}
