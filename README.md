# ASCII Art Web

## Description

A web application that converts input text to ASCII art using different banner styles.

</br>
  <section class="authors">
                <h1> AUTHORS :</h1> </br>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/fb9713e670165fd5fc7536ccc10a6c8d?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/tesiaka">tesiaka</a></p>
                </div>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/8d298bcc662dd253ab4426515673269d?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/jerootieno">jerootieno</a></p>
                </div>
                <div class="contributor">
                    <img src="https://learn.zone01kisumu.ke/git/avatars/cf0006d1b23256772956a4629c7a25a1?size=870">
                    <p><a href="https://learn.zone01kisumu.ke/git/viarony">viarony</a></p>
                </div>
            </section>

## Usage

To run the application:

1. Clone the repository:
    ```sh
    git clone https://github.com/jerootieno/ascii-art-web.git
    ```
2. Navigate to the project directory:
    ```sh
    cd ascii-art-web
    ```
3. Build and run the application:
    ```sh
    go build
    ./ascii-art-web
    ```
4. Open your web browser and navigate to `http://localhost:8081`.

## Implementation Details

### Algorithm

1. **Input Handling**:
    - The input text and banner style are received via an HTML form.
    - The text is split by newlines if it contains them.

2. **ASCII Art Generation**:
    - The application reads the specified banner file.
    - It converts each character of the input text to its ASCII art representation.
    - It validates that the input contains only printable ASCII characters.

3. **Error Handling**:
    - Returns appropriate HTTP status codes:
        - `200 OK`: If everything went without errors.
        - `404 Not Found`: If templates or banners are not found.
        - `400 Bad Request`: For incorrect requests.
        - `500 Internal Server Error`: For unhandled errors.
