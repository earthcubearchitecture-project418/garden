import shapely.geometry
import pyproj

# Set up projections
p_ll = pyproj.Proj(init='epsg:4326')
p_mt = pyproj.Proj(init='epsg:3857') # metric; same as EPSG:900913

# Create corners of rectangle to be transformed to a grid
nw = shapely.geometry.Point((-5.0, 40.0))
se = shapely.geometry.Point((-4.0, 41.0))

stepsize = 5000 # 5 km grid step size

# Project corners to target projection
s = pyproj.transform(p_ll, p_mt, nw.x, nw.y) # Transform NW point to 3857
e = pyproj.transform(p_ll, p_mt, se.x, se.y) # .. same for SE

# Iterate over 2D area
gridpoints = []
x = s[0]
while x < e[0]:
    y = s[1]
    while y < e[1]:
        p = shapely.geometry.Point(pyproj.transform(p_mt, p_ll, x, y))
        gridpoints.append(p)
        y += stepsize
    x += stepsize

with open('testout.csv', 'wb') as of:
    of.write('lon;lat\n')
    for p in gridpoints:
        of.write('{:f};{:f}\n'.format(p.x, p.y))