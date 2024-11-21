
// Track the first Enter key press
let isFirstEnterPress = true;

// Get the textarea and the info bubble element
const textarea = document.getElementById('text');
const infoBubble = document.getElementById('info-bubble');

// Event listener for Enter key press
textarea.addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        event.preventDefault();  // Prevent the default "Enter" behavior (which inserts a new line)

        // If this is the first Enter press, show the bubble
        if (isFirstEnterPress) {
            infoBubble.style.display = 'block'; // Show the bubble
            const rect = textarea.getBoundingClientRect();
            infoBubble.style.left = rect.left - infoBubble.offsetWidth - 10 + 'px'; // Position the bubble to the left
            infoBubble.style.top = rect.top + (rect.height / 2) - (infoBubble.offsetHeight / 2) + 'px'; // Center it vertically
            isFirstEnterPress = false;  // Set the flag to false after first Enter press

            // Hide the bubble after 3 seconds
            setTimeout(() => {
                infoBubble.style.display = 'none';
            }, 4000); // Bubble will disappear after 3 seconds
        }

        // Insert the string "\n" at the cursor position
        const cursorPos = this.selectionStart;
        const currentValue = this.value;
        this.value = currentValue.slice(0, cursorPos) + '\\n' + currentValue.slice(cursorPos);
        this.selectionStart = this.selectionEnd = cursorPos + 3; // Move cursor after the inserted \n
    }
});
window.addEventListener('load', function () {
    const [navigationEntry] = window.performance.getEntriesByType('navigation');
    
    if (navigationEntry && navigationEntry.type === 'reload') {
        window.location.href = "/"; // Redirect to the main page on refresh
    }
});
