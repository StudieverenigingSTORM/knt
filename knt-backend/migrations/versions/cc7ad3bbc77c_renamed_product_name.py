"""renamed product name

Revision ID: cc7ad3bbc77c
Revises: d23b2705375a
Create Date: 2022-01-08 00:44:36.987236

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'cc7ad3bbc77c'
down_revision = 'd23b2705375a'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('product', sa.Column('name', sa.Text(), nullable=True))
    op.drop_column('product', 'product_name')
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('product', sa.Column('product_name', sa.TEXT(), nullable=True))
    op.drop_column('product', 'name')
    # ### end Alembic commands ###
