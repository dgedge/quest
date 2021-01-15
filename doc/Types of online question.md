## A REST server in Go
Use gin

```sh
go get -u github.com/gin-gonic/gin
```



## Types of online question

Multiple choice
Multiple response
True of false
Short answer
numeric answer
Fill the gaps in text
matching left to right images
sorting images
selecting a hotspot
drag and drop
drag words
select from lists to fill gaps
Likert scale
essay

How can these very different types of question be represented, randomised and their correct and attempted answers be stored?

id:		 A guid for the question
subject
topic
level
type:   The type of question 
posed:  The text explaining the question
weight 

0..n RESPONSES
qu_id
text
randomised values
canvas draw svg

0..n col1 images 
0..n col2 images



Answer
qu_id
[text] short or essay answer of fill the gaps
float  numeric
[bool] cardinality of 1  for True False
[bool]  multiple choice/response
[int]   image order
[(x,y)] hot spots
[i, j]  column matches
  
## Randomization
This can greatly increase the number of questions available and is particulary suited to questions that generate an image on canvas.  



## Using the Canvas in React
[documentation](https://www.npmjs.com/package/react-canvas-draw)

[Great explanation here](https://medium.com/better-programming/add-an-html-canvas-into-your-react-app-176dab099a79)

## Rendering Mathematics
Use Mathjax in React as described [here](https://medium.com/@roopamg777/render-mathematical-equations-latex-using-mathjax-and-react-hooks-e69e36523fff)

```js
import React, { useEffect } from 'react';
function Latex(props) {
  let node = React.createRef();
  useEffect(() => {
    renderMath();
  });
  const renderMath = () => {
    window.MathJax.Hub.Queue([
      "Typeset", 
      window.MathJax.Hub,
      node.current
   ]);
  }
  return (
    <div ref={node}>
      {props.children}
    </div>
  );
}
export default Latex;
```

 For react-mathjax use this [repo](https://github.com/wko27/react-mathjax)
 Alternative is [Katex](https://levelup.gitconnected.com/adding-katex-and-markdown-in-react-7b70694004ef) which renders faster