function makeVisible() {
    const elements = document.getElementsByClassName("visible")
    for (let i = 0; i < elements.length; i++) {
        if (elements[i].style.visibility === "visible")
            elements[i].style.visibility = "hidden"
        else
            elements[i].style.visibility = "visible"
        
    }
    
}