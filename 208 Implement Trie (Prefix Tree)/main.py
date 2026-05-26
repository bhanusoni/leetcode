class TrieNode:
    def __init__(self):
        self.characters = {}
        self.end_of_word = False

class Trie:

    def __init__(self):
        self.node = TrieNode()

    def insert(self, word: str) -> None:
        curr = self.node
        for ch in word:
            if ch not in curr.characters:
                curr.characters[ch] = TrieNode()
            curr = curr.characters[ch]
        curr.end_of_word = True
        

    def search(self, word: str) -> bool:
        curr = self.node
        for ch in word:
            if ch not in curr.characters:
                return False
            curr = curr.characters[ch]
        return curr.end_of_word

    def startsWith(self, prefix: str) -> bool:
        curr = self.node
        for ch in prefix:
            if ch not in curr.characters:
                return False
            curr = curr.characters[ch]
        return True
        


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)