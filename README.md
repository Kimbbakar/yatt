## YATT

A light weight command line note taking tools.

#### Installation
```
go get github.com/Kimbbakar/yatt
```
#### Features from v1.0 release
1. Create note
    ```
    yatt create - n "my first note"
    ``` 

2. List notes
    ```
    yatt list [-t]

    # Response
    ID: 35410507ede64956a48ed4f5a09c10f8
    Date: Thu, 16 Dec 2021 10:37:45 +06

        Note: my 2nd note

    ID: 98ad7029b0da4862b973d2549d7967ea
    Date: Thu, 16 Dec 2021 10:37:38 +06

        Note: my first note

    ```

3. Delete Note
    ```
    yatt delete <id>
    ```

4. Flash/Remove all notes
    ```
    yatt flash
    ```