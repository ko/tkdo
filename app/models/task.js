import DS from 'ember-data';

let {
    Model,
    attr
} = DS;

export default Model.extend({
  title: attr('string'),
  author: attr('string'),
  dateCreated: attr('date'),
  dateCompleted: attr('date'),
  description: attr('string')
});
