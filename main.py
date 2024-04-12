# Name : Ritesh Koushik
# Roll : CB.EN.U4_CSE22038

def LFU_page_faults(n_pages, memory_capacity, page_references):
    page_faults = 0
    memory = []
    page_frequency = {}

    for page in page_references:
        if page not in memory:
            if len(memory) == memory_capacity:
                least_frequent_page = min(memory, key=page_frequency.get)
                memory.remove(least_frequent_page)
                del page_frequency[least_frequent_page]

            memory.append(page)
            page_faults += 1

        page_frequency[page] = page_frequency.get(page, 0) + 1
        print(f"After reference {page}: Memory = {memory}")

    return page_faults


def MFU_page_faults(n_pages, memory_capacity, page_references):
    page_faults = 0
    memory = []
    page_frequency = {}
    for page in page_references:
        if page not in memory:
            if len(memory) == memory_capacity:
                most_frequent_page = max(memory, key=page_frequency.get)
                memory.remove(most_frequent_page)
                del page_frequency[most_frequent_page]
                
            memory.append(page)
            page_faults += 1

        page_frequency[page] = page_frequency.get(page, 0) + 1
        print(f"After reference {page}: Memory = {memory}")
        return page_faults

# Driver code

if __name__ == '__main__':
    page_references = [1, 2, 3, 4, 2, 1, 5]
    n_pages = len(page_references)
    memory_capacity = 3
    print("LFU Algorithm:")
    lfu_faults = LFU_page_faults(n_pages, memory_capacity, page_references)
    print("Page Faults =", lfu_faults)
    print("Page Hits =", n_pages - lfu_faults)
    print("\nMFU Algorithm:")
    mfu_faults = MFU_page_faults(n_pages, memory_capacity, page_references)
    print("Page Faults =", mfu_faults)
    print("Page Hits =", n_pages - mfu_faults)
