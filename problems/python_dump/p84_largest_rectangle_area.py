def largest_rectangle_area(heights):
    stack = []
    lra = 0
    for i, h in enumerate(heights + [0]):
        while stack and heights[stack[-1]] > h:
            height = heights[stack.pop()]
            width = i if not stack else i - stack[-1] - 1
            lra = max(lra, height * width)
        stack.append(i)
    return lra
