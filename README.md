# Moby Dick

## The Project
Using the text document (`mobydick.txt`), find the 100 most used words that does not include any of the stop words (`stop-words.txt`)
  

## Challenges
The biggest challenge by far is the fact that there's multiple different kinds of hyphens and apostrophes in the text that's being read differently depending on the language.

I planned on having PHP, Java, as well as C#; but because I'm only getting a couple hours a day to work on it due to my current work schedule I was only able to get the scaffolding for PHP and Java up which were only stored locally. PHP as well was reading the special characters much differently which is making just streaming in `mobydick.txt` very challenging because any modifications to the resulting string causes the unknown character symbols to appear.

I did also find one more thing...the results. In my editor a simple search produces ~1.1k results for whale, while the program finds ~1.2k; however, searching Google says the word whale occurs 1,685 times. I'm not sure if this is just variations in the the exact copy of the book or what, but this did prove to be difficult because it's hard to make sure your solution is correct when you don't know the answer. I ended up finding words with a lower amount of uses and comparing that which helped me find multiple issues, like with the hyphens.

## Branches
I wrote the Javascript/HTML branches to be able to be run easily locally. I personally used [http-server](https://www.npmjs.com/package/http-server) for testing/developing it. The only branch that requires more is the Go branch because it requires Go installed locally, but it doesn't require any other modules besides what's packaged with the language so there's no more setup required.

### [HTML](https://github.com/MatthewSH/sa-mobydick/tree/html)
This branch just contains a HTML document that uses jQuery for DOM manipulation and event handling. Just provides very basic functionality.
It uses [Tailwind](https://tailwindcss.com/) for simple styling.

### [VueJS](https://github.com/MatthewSH/sa-mobydick/tree/vuejs)
VueJS is something I use more frequently, so rewriting the HTML branch to VueJS was much easier for me. It uses the same searching and counting methods, but implemented for VueJS. 
I use Axios for web requests, Tailwind for styling, and [Animate](https://daneden.github.io/animate.css/) for basic animations.
This branch is different because on top of the basic functionality (analyze and spit out results in a long list) I also provide a 3 column style to prevent the need to scroll as well as a custom pagination-style page switcher that will allow you to view more than just the top 100 results. It also provides a loading screen, which isn't provided by the HTML branch.

### [Go](https://github.com/MatthewSH/sa-mobydick/tree/go)
The Go implementation was just really to challenge myself. I haven't worked with Go before except for [one other time](https://github.com/BultApp/Ember)  2 years ago because I needed a simple and efficient way to handle downloading and unzipping files, I only wrote that because NodeJS (including the packages available on NPM) didn't provide a stable or functioning solution at the time that worked in the environment and way I needed it to work.
Writing this code actually showed me a lot about go, specifically how it handles lists which is much differently then I've ever seen in another language. Mainly that it doesn't track the order, just the content.

I wasn't able to provide a UI element for this one, it just spits out the results at runtime in the console.